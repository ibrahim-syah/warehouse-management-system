package usecase

import (
	"context"
	"errors"
	"fmt"
	"warehouse-management-system/entity"
	"warehouse-management-system/repo"
	"warehouse-management-system/sentinel"
)

type ProductUsecase interface {
	AddProduct(ctx context.Context, product *entity.Product) (int, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, productID int) error
	GetProductByID(ctx context.Context, productID int) (*entity.Product, error)
	GetProducts(ctx context.Context, req *entity.PaginationParam) ([]entity.Product, int, error)
}

type productUsecase struct {
	transactor   repo.Transactor
	productRepo  repo.ProductRepo
	locationRepo repo.LocationRepo
}

func NewProductUsecase(
	transactor repo.Transactor,
	productRepo repo.ProductRepo,
	locationRepo repo.LocationRepo,
) ProductUsecase {
	return &productUsecase{
		transactor:   transactor,
		productRepo:  productRepo,
		locationRepo: locationRepo,
	}
}

func (u *productUsecase) AddProduct(ctx context.Context, product *entity.Product) (int, error) {
	var productID int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		item, txErr := u.productRepo.GetProductBySKU(txCtx, product.SKU)
		if txErr != nil && !errors.Is(txErr, sentinel.ErrNotFound) {
			return txErr
		}
		if item != nil {
			txErr = fmt.Errorf("%w: item with the given sku already exist", sentinel.ErrUsecaseError)
			return txErr
		}

		location, txErr := u.locationRepo.GetLocationByID(txCtx, product.LocationID)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = fmt.Errorf("%w: location not found", sentinel.ErrUsecaseError)
			}
			return txErr
		}

		totalQuantity, txErr := u.productRepo.GetTotalQuantityByLocationID(txCtx, product.LocationID)
		if txErr != nil {
			return txErr
		}

		if totalQuantity+product.Quantity > location.Capacity {
			txErr = fmt.Errorf("%w: not enough capacity in the warehouse", sentinel.ErrUsecaseError)
			return txErr
		}

		id, txErr := u.productRepo.InsertProduct(txCtx, product)
		if txErr != nil {
			return txErr
		}
		productID = id
		return nil
	})
	if err != nil {
		return -1, err
	}

	return productID, nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		location, txErr := u.locationRepo.GetLocationByID(txCtx, product.LocationID)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = fmt.Errorf("%w: location not found", sentinel.ErrUsecaseError)
			}
			return txErr
		}

		totalQuantity, txErr := u.productRepo.GetTotalQuantityByLocationID(txCtx, product.LocationID)
		if txErr != nil {
			return txErr
		}

		existingProduct, txErr := u.productRepo.GetProductByID(txCtx, product.ID)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = fmt.Errorf("%w: product not found", sentinel.ErrNotFound)
			}
			return txErr
		}

		if totalQuantity-existingProduct.Quantity+product.Quantity > location.Capacity {
			txErr = fmt.Errorf("%w: not enough capacity in the warehouse", sentinel.ErrUsecaseError)
			return txErr
		}

		return u.productRepo.UpdateProduct(txCtx, product)
	})
}

func (u *productUsecase) DeleteProduct(ctx context.Context, productID int) error {
	return u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		return u.productRepo.DeleteProduct(txCtx, productID)
	})
}

func (u *productUsecase) GetProductByID(ctx context.Context, productID int) (*entity.Product, error) {
	var product *entity.Product
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		var txErr error
		product, txErr = u.productRepo.GetProductByID(txCtx, productID)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *productUsecase) GetProducts(ctx context.Context, req *entity.PaginationParam) ([]entity.Product, int, error) {
	var products []entity.Product
	var count int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		txProducts, txCount, txErr := u.productRepo.GetProducts(txCtx, req)
		if txErr != nil {
			return txErr
		}
		products = txProducts
		count = txCount
		return nil
	})
	if err != nil {
		return nil, -1, err
	}

	return products, count, nil
}
