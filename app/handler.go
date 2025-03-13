package app

import (
	"warehouse-management-system/handler"
)

type appHandlers struct {
	authHandler    handler.AuthHandler
	productHandler handler.ProductHandler
}

func SetupHandler(usecases *appUsecases) *appHandlers {
	authHandler := handler.NewAuthHandler(usecases.authUsecase)
	productHandler := handler.NewProductHandler(usecases.productUsecase)

	return &appHandlers{
		authHandler:    authHandler,
		productHandler: productHandler,
	}
}
