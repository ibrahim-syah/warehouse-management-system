package usecase

import (
	"context"
	"errors"
	"warehouse-management-system/entity"
	"warehouse-management-system/repo"
	"warehouse-management-system/sentinel"
)

type UserUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUsers(ctx context.Context, req *entity.PaginationParam) ([]entity.User, int, error)
}

type userUsecase struct {
	transactor repo.Transactor
	userRepo   repo.UserRepo
}

func NewUserUsecase(
	transactor repo.Transactor,
	userRepo repo.UserRepo,
) UserUsecase {
	return &userUsecase{
		transactor: transactor,
		userRepo:   userRepo,
	}
}

func (u *userUsecase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		txUser, txErr := u.userRepo.GetUserByEmail(txCtx, email)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = sentinel.ErrInvalidCredential
			}
			return txErr
		}

		user = txUser
		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetUsers(ctx context.Context, req *entity.PaginationParam) ([]entity.User, int, error) {
	var user []entity.User
	var count int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		txItems, txItemsCount, txErr := u.userRepo.GetUsers(txCtx, req)
		if txErr != nil {
			return txErr
		}

		user = txItems
		count = txItemsCount
		return nil
	})
	if err != nil {
		return nil, -1, err
	}

	return user, count, nil
}
