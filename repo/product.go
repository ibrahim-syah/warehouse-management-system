package repo

import (
	"database/sql"
)

type ProductRepo interface {
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
