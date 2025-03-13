package sentinel

import "errors"

var (
	ErrAlreadyExist       = errors.New("data already exist")
	ErrNotFound           = errors.New("data not found")
	ErrInvalidCredential  = errors.New("invalid credential")
	ErrInvalidInput       = errors.New("invalid input")
	ErrUsecaseError       = errors.New("usecase error")
	ErrUnauthorized       = errors.New("unauthorized error")
	ErrInvalidTransaction = errors.New("the given context does not have an ongoing transaction")
)
