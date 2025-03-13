package app

import (
	"warehouse-management-system/handler"
)

type appHandlers struct {
	authHandler    handler.AuthHandler
	userHandler    handler.UserHandler
	productHandler handler.ProductHandler
}

func SetupHandler(usecases *appUsecases) *appHandlers {
	authHandler := handler.NewAuthHandler(usecases.authUsecase)
	userHandler := handler.NewUserHandler(usecases.userUsecase)
	productHandler := handler.NewProductHandler(usecases.productUsecase)

	return &appHandlers{
		authHandler:    authHandler,
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}
