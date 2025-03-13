package usecase

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"warehouse-management-system/entity"
	"warehouse-management-system/repo"
	"warehouse-management-system/sentinel"
)

type OrderUsecase interface {
	ProcessOrder(ctx context.Context, order *entity.Order) (int, error)
}

type orderUsecase struct {
	transactor   repo.Transactor
	orderRepo    repo.OrderRepo
	productRepo  repo.ProductRepo
	locationRepo repo.LocationRepo
	mu           sync.Mutex // Mutex for synchronizing order processing
}

func NewOrderUsecase(
	transactor repo.Transactor,
	orderRepo repo.OrderRepo,
	productRepo repo.ProductRepo,
	locationRepo repo.LocationRepo,
) OrderUsecase {
	return &orderUsecase{
		transactor:   transactor,
		orderRepo:    orderRepo,
		productRepo:  productRepo,
		locationRepo: locationRepo,
	}
}

func (u *orderUsecase) ProcessOrder(ctx context.Context, order *entity.Order) (int, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	var orderID int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		product, txErr := u.productRepo.GetProductByID(txCtx, order.ProductID)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = fmt.Errorf("%w: product not found", sentinel.ErrUsecaseError)
			}
			return txErr
		}

		location, txErr := u.locationRepo.GetLocationByID(txCtx, product.LocationID)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = fmt.Errorf("%w: location not found", sentinel.ErrUsecaseError)
			}
			return txErr
		}

		if txErr := u.validateOrder(txCtx, order, product, location); txErr != nil {
			return txErr
		}

		id, txErr := u.orderRepo.InsertOrder(txCtx, order)
		if txErr != nil {
			return fmt.Errorf("failed to insert order: %w", txErr)
		}
		orderID = id

		done := make(chan error)
		go func() {
			// Simulate long processing
			time.Sleep(2 * time.Second)

			if order.Type == repo.OrderTypeShip {
				product.Quantity -= order.Quantity
			} else {
				product.Quantity += order.Quantity
			}

			if txErr := u.productRepo.UpdateProduct(txCtx, product); txErr != nil {
				done <- fmt.Errorf("failed to update product: %w", txErr)
				return
			}

			done <- nil
		}()

		// Wait for processing to complete within transaction
		if txErr := <-done; txErr != nil {
			return txErr
		}

		return nil
	})

	if err != nil {
		return -1, err
	}

	return orderID, nil
}

func (u *orderUsecase) validateOrder(ctx context.Context, order *entity.Order, product *entity.Product, location *entity.Location) error {

	if order.Type == repo.OrderTypeShip && product.Quantity < order.Quantity {
		return fmt.Errorf("%w: insufficient product quantity", sentinel.ErrUsecaseError)
	}

	totalQuantity, err := u.productRepo.GetTotalQuantityByLocationID(ctx, location.ID)
	if err != nil {
		return fmt.Errorf("failed to get total quantity: %w", err)
	}

	if order.Type == repo.OrderTypeReceive {
		if totalQuantity+order.Quantity > location.Capacity {
			return fmt.Errorf("%w: warehouse capacity exceeded", sentinel.ErrUsecaseError)
		}
	}

	return nil
}
