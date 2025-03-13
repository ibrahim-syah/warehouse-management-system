package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"warehouse-management-system/entity"
	"warehouse-management-system/sentinel"
)

const (
	OrderTypeReceive = "receive"
	OrderTypeShip    = "ship"
)

type OrderRepo interface {
	InsertOrder(ctx context.Context, order *entity.Order) (int, error)
	GetOrderByID(ctx context.Context, orderID int) (*entity.Order, error)
	GetOrders(ctx context.Context, req *entity.PaginationParam) ([]entity.Order, int, error)
}

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) InsertOrder(ctx context.Context, order *entity.Order) (int, error) {
	query := `
    INSERT INTO orders (product_id, quantity, type, created_at, updated_at)
    VALUES ($1, $2, $3, NOW(), NOW())
    RETURNING id`

	var ID int
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, order.ProductID, order.Quantity, order.Type).Scan(&ID)
	} else {
		err = r.db.QueryRowContext(ctx, query, order.ProductID, order.Quantity, order.Type).Scan(&ID)
	}
	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (r *orderRepo) GetOrderByID(ctx context.Context, orderID int) (*entity.Order, error) {
	query := `
    SELECT id, product_id, quantity, type, created_at, updated_at, deleted_at
    FROM orders
    WHERE id = $1`

	var order entity.Order
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, orderID).Scan(
			&order.ID, &order.ProductID, &order.Quantity, &order.Type,
			&order.CreatedAt, &order.UpdatedAt, &order.DeletedAt,
		)
	} else {
		err = r.db.QueryRowContext(ctx, query, orderID).Scan(
			&order.ID, &order.ProductID, &order.Quantity, &order.Type,
			&order.CreatedAt, &order.UpdatedAt, &order.DeletedAt,
		)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = sentinel.ErrNotFound
		}
		return nil, err
	}

	return &order, nil
}

func (r *orderRepo) GetOrders(ctx context.Context, req *entity.PaginationParam) ([]entity.Order, int, error) {
	query := `
    SELECT id, product_id, quantity, type, created_at, updated_at, deleted_at
    FROM orders
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

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(
			&order.ID, &order.ProductID, &order.Quantity, &order.Type,
			&order.CreatedAt, &order.UpdatedAt, &order.DeletedAt,
		)
		if err != nil {
			return nil, -1, err
		}
		orders = append(orders, order)
	}

	query = `SELECT COUNT(id) FROM orders`
	var count int
	if tx != nil {
		err = tx.QueryRowContext(ctx, query).Scan(&count)
	} else {
		err = r.db.QueryRowContext(ctx, query).Scan(&count)
	}
	if err != nil {
		return nil, -1, err
	}

	return orders, count, nil
}
