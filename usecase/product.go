package usecase

import (
	"warehouse-management-system/repo"
)

type ProductUsecase interface {
}

type productUsecase struct {
	transactor repo.Transactor
}

func NewProductUsecase(
	transactor repo.Transactor,
) ProductUsecase {
	return &productUsecase{
		transactor: transactor,
	}
}
