package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"warehouse-management-system/entity"
	"warehouse-management-system/sentinel"
)

type LocationRepo interface {
	InsertLocation(ctx context.Context, location *entity.Location) (int, error)
	GetLocations(ctx context.Context, req *entity.PaginationParam) ([]entity.Location, int, error)
	GetLocationByID(ctx context.Context, locationID int) (*entity.Location, error)
}

type locationRepo struct {
	db *sql.DB
}

func NewLocationRepo(db *sql.DB) LocationRepo {
	return &locationRepo{
		db: db,
	}
}

func (r *locationRepo) InsertLocation(ctx context.Context, location *entity.Location) (int, error) {
	query := `
    INSERT INTO warehouse_locations (name, capacity, created_at, updated_at)
    VALUES ($1, $2, NOW(), NOW())
    RETURNING id`

	var ID int

	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, location.Name, location.Capacity).Scan(&ID)
	} else {
		err = r.db.QueryRowContext(ctx, query, location.Name, location.Capacity).Scan(&ID)
	}
	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (r *locationRepo) GetLocations(ctx context.Context, req *entity.PaginationParam) ([]entity.Location, int, error) {
	query := `
    SELECT id, name, capacity, created_at, updated_at, deleted_at
    FROM warehouse_locations
    ORDER BY %s %s
    LIMIT $1 OFFSET $2`
	query = fmt.Sprintf(query, req.OrderBy, req.OrderDirection)

	tx := extractTx(ctx)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, req.Limit, req.Offset)
	} else {
		rows, err = r.db.QueryContext(ctx, query, req.Limit, req.Offset)
	}
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var locations []entity.Location
	for rows.Next() {
		var location entity.Location
		err := rows.Scan(&location.ID, &location.Name, &location.Capacity, &location.CreatedAt, &location.UpdatedAt, &location.DeletedAt)
		if err != nil {
			return nil, -1, err
		}
		locations = append(locations, location)
	}

	query = `SELECT COUNT(id) FROM warehouse_locations`
	var count int
	err = r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return nil, -1, err
	}

	return locations, count, nil
}

func (r *locationRepo) GetLocationByID(ctx context.Context, locationID int) (*entity.Location, error) {
	query := `
	SELECT id, name, capacity, created_at, updated_at, deleted_at
	FROM warehouse_locations
	WHERE id = $1`

	var location entity.Location
	var err error
	tx := extractTx(ctx)
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, locationID).Scan(
			&location.ID, &location.Name, &location.Capacity, &location.CreatedAt, &location.UpdatedAt, &location.DeletedAt,
		)
	} else {
		err = r.db.QueryRowContext(ctx, query, locationID).Scan(
			&location.ID, &location.Name, &location.Capacity, &location.CreatedAt, &location.UpdatedAt, &location.DeletedAt,
		)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = sentinel.ErrNotFound
		}
		return nil, err
	}

	return &location, nil
}
