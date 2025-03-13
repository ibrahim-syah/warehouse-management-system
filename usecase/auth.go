package usecase

import (
	"context"
	"errors"
	"fmt"
	"warehouse-management-system/entity"
	"warehouse-management-system/repo"
	"warehouse-management-system/sentinel"
	jwtutils "warehouse-management-system/utils/jwt"
	passwordutils "warehouse-management-system/utils/password"
)

type AuthUsecase interface {
	Login(ctx context.Context, req *entity.EmailPassword) (*entity.LoginToken, error)
	Register(ctx context.Context, req *entity.InsertUser) (int, error)
}

type authUsecase struct {
	transactor repo.Transactor
	userRepo   repo.UserRepo
}

func NewAuthUsecase(
	transactor repo.Transactor,
	userRepo repo.UserRepo,
) AuthUsecase {
	return &authUsecase{
		transactor: transactor,
		userRepo:   userRepo,
	}
}

func (u *authUsecase) Login(ctx context.Context, req *entity.EmailPassword) (*entity.LoginToken, error) {

	var user entity.User
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		txUser, txErr := u.userRepo.GetUserByEmail(txCtx, req.Email)
		if txErr != nil {
			if errors.Is(txErr, sentinel.ErrNotFound) {
				txErr = sentinel.ErrInvalidCredential
			}
			return txErr
		}

		user = *txUser
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = passwordutils.VerifyPasswordHash(user.Password, req.Password)
	if err != nil {
		err = sentinel.ErrInvalidCredential
		return nil, err
	}

	clientClaims := jwtutils.CustomClaims{
		ClientID: user.ID,
		Email:    user.Email,
		Role:     user.Role,
	}
	jwt, err := jwtutils.GenerateJWT(clientClaims)
	if err != nil {
		return nil, err
	}

	res := &entity.LoginToken{
		UserID:      user.ID,
		AccessToken: jwt,
	}
	return res, nil
}

func (u *authUsecase) Register(ctx context.Context, req *entity.InsertUser) (int, error) {

	var userID int
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		user, txErr := u.userRepo.GetUserByEmail(txCtx, req.Email)
		if txErr != nil && !errors.Is(txErr, sentinel.ErrNotFound) {
			return txErr
		}
		if user != nil {
			txErr = fmt.Errorf("%w: user with the given email already exist", sentinel.ErrUsecaseError)
			return txErr
		}

		hashedPassword, txErr := passwordutils.GeneratePasswordHash(req.Password)
		if txErr != nil {
			return txErr
		}
		req.Password = hashedPassword

		id, txErr := u.userRepo.InsertUser(txCtx, req)
		if txErr != nil {
			return txErr
		}
		userID = id

		return nil
	})
	if err != nil {
		return -1, err
	}

	return userID, nil
}
