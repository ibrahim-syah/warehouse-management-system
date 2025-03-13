package app

import (
	"database/sql"
	"warehouse-management-system/repo"
)

type appRepositories struct {
	transactor   repo.Transactor
	userRepo     repo.UserRepo
	productRepo  repo.ProductRepo
	locationRepo repo.LocationRepo
}

func SetupRepositories(db *sql.DB) *appRepositories {
	return &appRepositories{
		transactor:   repo.NewTransactor(db),
		productRepo:  repo.NewProductRepo(db),
		userRepo:     repo.NewUserRepo(db),
		locationRepo: repo.NewLocationRepo(db),
	}
}
