package app

import (
	"warehouse-management-system/handler"
)

type appHandlers struct {
	productHandler handler.ProductHandler
}

func SetupHandler(usecases *appUsecases) *appHandlers {
	productHandler := handler.NewProductHandler(usecases.productUsecase)

	return &appHandlers{
		productHandler: productHandler,
	}
}
