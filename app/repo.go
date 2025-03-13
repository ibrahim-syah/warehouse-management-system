package app

import (
	"database/sql"
	"warehouse-management-system/repo"
)

type appRepositories struct {
	transactor  repo.Transactor
	productRepo repo.ProductRepo
}

func SetupRepositories(db *sql.DB) *appRepositories {
	return &appRepositories{
		transactor:  repo.NewTransactor(db),
		productRepo: repo.NewProductRepo(db),
	}
}
