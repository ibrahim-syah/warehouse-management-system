package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"warehouse-management-system/dto"
	"warehouse-management-system/entity"
	"warehouse-management-system/sentinel"
	"warehouse-management-system/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(
	authUsecase usecase.AuthUsecase,
) AuthHandler {
	handler := &authHandler{
		authUsecase: authUsecase,
	}

	return handler
}

func (h *authHandler) Login(ctx *gin.Context) {
	req := dto.LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ute *json.UnmarshalTypeError
		var ve validator.ValidationErrors
		if errors.As(err, &ute) || errors.As(err, &ve) {
			ctx.Error(err)
			return
		}
		err = sentinel.ErrInvalidInput
		ctx.Error(err)
		return
	}

	credential := entity.EmailPassword{Email: req.Email, Password: req.Password}
	loginToken, err := h.authUsecase.Login(ctx.Request.Context(), &credential)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := dto.LoginResponse{
		ID:          loginToken.UserID,
		AccessToken: loginToken.AccessToken,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "login success",
		Data:    res,
	})
}

func (h *authHandler) Register(ctx *gin.Context) {
	req := dto.RegisterRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ute *json.UnmarshalTypeError
		var ve validator.ValidationErrors
		if errors.As(err, &ute) || errors.As(err, &ve) {
			ctx.Error(err)
			return
		}
		err = sentinel.ErrInvalidInput
		ctx.Error(err)
		return
	}

	newUser := entity.InsertUser{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
	_, err := h.authUsecase.Register(ctx.Request.Context(), &newUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: "registered successfully",
	})
}
