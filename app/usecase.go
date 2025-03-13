package app

import "warehouse-management-system/usecase"

type appUsecases struct {
	productUsecase usecase.ProductUsecase
}

func SetupUsecases(repositories *appRepositories) *appUsecases {

	productUsecase := usecase.NewProductUsecase(
		repositories.transactor,
	)

	return &appUsecases{
		productUsecase: productUsecase,
	}
}
