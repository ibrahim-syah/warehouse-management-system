package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"warehouse-management-system/entity"
	"warehouse-management-system/sentinel"
)

type ProductRepo interface {
	InsertProduct(ctx context.Context, product *entity.Product) (int, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, productID int) error
	GetProductByID(ctx context.Context, productID int) (*entity.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (*entity.Product, error)
	GetProducts(ctx context.Context, req *entity.PaginationParam) ([]entity.Product, int, error)
	GetTotalQuantityByLocationID(ctx context.Context, locationID int) (int, error)
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) InsertProduct(ctx context.Context, product *entity.Product) (int, error) {
	query := `
    INSERT INTO products (name, sku, quantity, location_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, NOW(), NOW())
    RETURNING id`

	var ID int
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, product.Name, product.SKU, product.Quantity, product.LocationID).Scan(&ID)
	} else {
		err = r.db.QueryRowContext(ctx, query, product.Name, product.SKU, product.Quantity, product.LocationID).Scan(&ID)
	}
	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, product *entity.Product) error {
	query := `
    UPDATE products
    SET name = $1, sku = $2, quantity = $3, location_id = $4, updated_at = NOW()
    WHERE id = $5`

	tx := extractTx(ctx)
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, product.Name, product.SKU, product.Quantity, product.LocationID, product.ID)
	} else {
		_, err = r.db.ExecContext(ctx, query, product.Name, product.SKU, product.Quantity, product.LocationID, product.ID)
	}
	return err
}

func (r *productRepo) DeleteProduct(ctx context.Context, productID int) error {
	query := `DELETE FROM products WHERE id = $1`
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, productID)
	} else {
		_, err = r.db.ExecContext(ctx, query, productID)
	}
	return err
}

func (r *productRepo) GetProductByID(ctx context.Context, productID int) (*entity.Product, error) {
	query := `
    SELECT id, name, sku, quantity, location_id, created_at, updated_at, deleted_at
    FROM products
    WHERE id = $1`

	var product entity.Product
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, productID).Scan(
			&product.ID, &product.Name, &product.SKU, &product.Quantity, &product.LocationID,
			&product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
		)
	} else {
		err = r.db.QueryRowContext(ctx, query, productID).Scan(
			&product.ID, &product.Name, &product.SKU, &product.Quantity, &product.LocationID,
			&product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
		)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = sentinel.ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetProductBySKU(ctx context.Context, sku string) (*entity.Product, error) {
	query := `
	SELECT id, name, sku, quantity, location_id, created_at, updated_at, deleted_at
	FROM products
	WHERE sku = $1`

	var product entity.Product
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, sku).Scan(
			&product.ID, &product.Name, &product.SKU, &product.Quantity, &product.LocationID,
			&product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
		)
	} else {
		err = r.db.QueryRowContext(ctx, query, sku).Scan(
			&product.ID, &product.Name, &product.SKU, &product.Quantity, &product.LocationID,
			&product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
		)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = sentinel.ErrNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) GetProducts(ctx context.Context, req *entity.PaginationParam) ([]entity.Product, int, error) {
	query := `
    SELECT id, name, sku, quantity, location_id, created_at, updated_at, deleted_at
    FROM products
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

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.SKU, &product.Quantity, &product.LocationID,
			&product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return nil, -1, err
		}
		products = append(products, product)
	}

	query = `SELECT COUNT(id) FROM products`
	var count int
	if tx != nil {
		err = tx.QueryRowContext(ctx, query).Scan(&count)
	} else {
		err = r.db.QueryRowContext(ctx, query).Scan(&count)
	}
	if err != nil {
		return nil, -1, err
	}

	return products, count, nil
}

func (r *productRepo) GetTotalQuantityByLocationID(ctx context.Context, locationID int) (int, error) {
	query := `SELECT COALESCE(SUM(quantity), 0) FROM products WHERE location_id = $1`
	var totalQuantity int
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, locationID).Scan(&totalQuantity)
	} else {
		err = r.db.QueryRowContext(ctx, query, locationID).Scan(&totalQuantity)
	}
	if err != nil {
		return -1, err
	}

	return totalQuantity, nil
}
