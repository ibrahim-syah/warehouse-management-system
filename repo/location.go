package repo

import (
	"context"
	"database/sql"
	"fmt"
	"warehouse-management-system/entity"
)

type LocationRepo interface {
	InsertLocation(ctx context.Context, location *entity.Location) (int, error)
	GetLocations(ctx context.Context, req *entity.PaginationParam) ([]entity.Location, int, error)
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
	err := r.db.QueryRowContext(ctx, query, location.Name, location.Capacity).Scan(&ID)
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

	rows, err := r.db.QueryContext(ctx, query, req.Limit, req.Offset)
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
