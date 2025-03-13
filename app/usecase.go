package app

import "warehouse-management-system/usecase"

type appUsecases struct {
	authUsecase    usecase.AuthUsecase
	productUsecase usecase.ProductUsecase
}

func SetupUsecases(repositories *appRepositories) *appUsecases {

	authUsecase := usecase.NewAuthUsecase(
		repositories.transactor,
		repositories.userRepo,
	)
	productUsecase := usecase.NewProductUsecase(
		repositories.transactor,
	)

	return &appUsecases{
		authUsecase:    authUsecase,
		productUsecase: productUsecase,
	}
}
