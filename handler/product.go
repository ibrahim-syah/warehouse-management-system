package handler

import (
	"warehouse-management-system/usecase"
)

type ProductHandler interface {
}

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(
	productUsecase usecase.ProductUsecase,
) ProductHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}
