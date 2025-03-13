package usecase

import (
	"context"
	"warehouse-management-system/entity"
	"warehouse-management-system/repo"
)

type LocationUsecase interface {
	AddLocation(ctx context.Context, location *entity.Location) (int, error)
	GetLocations(ctx context.Context, req *entity.PaginationParam) ([]entity.Location, int, error)
}

type locationUsecase struct {
	transactor   repo.Transactor
	locationRepo repo.LocationRepo
}

func NewLocationUsecase(
	transactor repo.Transactor,
	locationRepo repo.LocationRepo,
) LocationUsecase {
	return &locationUsecase{
		transactor:   transactor,
		locationRepo: locationRepo,
	}
}

func (u *locationUsecase) AddLocation(ctx context.Context, location *entity.Location) (int, error) {
	var locationID int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		id, txErr := u.locationRepo.InsertLocation(txCtx, location)
		if txErr != nil {
			return txErr
		}
		locationID = id
		return nil
	})
	if err != nil {
		return -1, err
	}

	return locationID, nil
}

func (u *locationUsecase) GetLocations(ctx context.Context, req *entity.PaginationParam) ([]entity.Location, int, error) {
	var locations []entity.Location
	var count int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		txLocations, txCount, txErr := u.locationRepo.GetLocations(txCtx, req)
		if txErr != nil {
			return txErr
		}
		locations = txLocations
		count = txCount
		return nil
	})
	if err != nil {
		return nil, -1, err
	}

	return locations, count, nil
}
