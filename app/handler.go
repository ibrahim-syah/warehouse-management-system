package app

import (
	"warehouse-management-system/handler"
)

type appHandlers struct {
	authHandler     handler.AuthHandler
	userHandler     handler.UserHandler
	productHandler  handler.ProductHandler
	locationHandler handler.LocationHandler
	orderHandler    handler.OrderHandler
}

func SetupHandler(usecases *appUsecases) *appHandlers {
	authHandler := handler.NewAuthHandler(usecases.authUsecase)
	userHandler := handler.NewUserHandler(usecases.userUsecase)
	productHandler := handler.NewProductHandler(usecases.productUsecase)
	locationHandler := handler.NewLocationHandler(usecases.locationUsecase)
	orderHandler := handler.NewOrderHandler(usecases.orderUsecase)

	return &appHandlers{
		authHandler:     authHandler,
		userHandler:     userHandler,
		productHandler:  productHandler,
		locationHandler: locationHandler,
		orderHandler:    orderHandler,
	}
}
