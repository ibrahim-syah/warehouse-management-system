package handler

import (
	"fmt"
	"net/http"
	"warehouse-management-system/dto"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetCurrentUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

type userHandler struct {
	// userUsecase usecase.AuthUsecase
}

func NewUserHandler(
// userUsecase usecase.AuthUsecase,
) UserHandler {
	handler := &userHandler{
		// authUsecase: authUsecase,
	}

	return handler
}

func (h *userHandler) GetCurrentUser(ctx *gin.Context) {
	userID := ctx.GetInt("UserID")
	role := ctx.GetString("Role")

	// loginToken, err := h.authUsecase.Login(ctx.Request.Context(), &credential)
	// if err != nil {
	// 	ctx.Error(err)
	// 	return
	// }
	res := dto.LoginResponse{
		AccessToken: role,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: fmt.Sprintf("%v", userID),
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
