package handler

import (
	"net/http"
	"warehouse-management-system/dto"
	"warehouse-management-system/entity"
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
	req.DefaultIfEmpty()

	paginatorParam := &entity.PaginationParam{
		OrderBy:        req.SortedBy,
		OrderDirection: req.Sort,
		Limit:          req.Limit,
		Offset:         (req.Page - 1) * req.Limit,
	}

	users, count, err := h.userUsecase.GetUsers(ctx.Request.Context(), paginatorParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := []dto.UserItem{}
	for _, u := range users {
		item := dto.UserItem{
			ID:        u.ID,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
		}
		res = append(res, item)
	}

	paginator := dto.MappingPaginator(req.Page, req.Limit, count)
	ctx.JSON(http.StatusCreated, dto.Response{
		Message:   "get users success",
		Data:      res,
		Paginator: &paginator,
	})
}
