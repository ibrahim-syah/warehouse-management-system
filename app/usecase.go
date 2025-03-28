package app

import "warehouse-management-system/usecase"

type appUsecases struct {
	authUsecase     usecase.AuthUsecase
	userUsecase     usecase.UserUsecase
	productUsecase  usecase.ProductUsecase
	locationUsecase usecase.LocationUsecase
	orderUsecase    usecase.OrderUsecase
}

func SetupUsecases(repositories *appRepositories) *appUsecases {

	authUsecase := usecase.NewAuthUsecase(
		repositories.transactor,
		repositories.userRepo,
	)

	userUsecase := usecase.NewUserUsecase(
		repositories.transactor,
		repositories.userRepo,
	)

	productUsecase := usecase.NewProductUsecase(
		repositories.transactor,
		repositories.productRepo,
		repositories.locationRepo,
	)

	locationUsecase := usecase.NewLocationUsecase(
		repositories.transactor,
		repositories.locationRepo,
	)

	orderUsecase := usecase.NewOrderUsecase(
		repositories.transactor,
		repositories.orderRepo,
		repositories.productRepo,
		repositories.locationRepo,
	)

	return &appUsecases{
		authUsecase:     authUsecase,
		productUsecase:  productUsecase,
		userUsecase:     userUsecase,
		locationUsecase: locationUsecase,
		orderUsecase:    orderUsecase,
	}
}
