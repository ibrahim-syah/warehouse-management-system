package handler

import (
	"net/http"
	"warehouse-management-system/dto"
	"warehouse-management-system/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetCurrentUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) UserHandler {
	handler := &userHandler{
		userUsecase: userUsecase,
	}

	return handler
}

func (h *userHandler) GetCurrentUser(ctx *gin.Context) {
	email := ctx.GetString("Email")

	user, err := h.userUsecase.GetUserByEmail(ctx.Request.Context(), email)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := dto.UserItem{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "success",
		Data:    res,
	})
}

func (h *userHandler) GetUsers(ctx *gin.Context) {
	req := dto.GetUsersRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: "get users success",
		Data:    req,
	})
}
