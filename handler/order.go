package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"warehouse-management-system/dto"
	"warehouse-management-system/entity"
	"warehouse-management-system/sentinel"
	"warehouse-management-system/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderHandler interface {
	ProcessShippingOrder(ctx *gin.Context)
	ProcessReceiveOrder(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
}

type orderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(
	orderUsecase usecase.OrderUsecase,
) OrderHandler {
	handler := &orderHandler{
		orderUsecase: orderUsecase,
	}

	return handler
}

func (h *orderHandler) ProcessShippingOrder(ctx *gin.Context) {
	req := dto.ProcessOrderRequest{}
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

	order := &entity.Order{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Type:      "ship",
	}

	orderID, err := h.orderUsecase.ProcessOrder(ctx.Request.Context(), order)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := dto.ProcessOrderResponse{
		OrderID: orderID,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "order processed",
		Data:    res,
	})
}

func (h *orderHandler) ProcessReceiveOrder(ctx *gin.Context) {
	req := dto.ProcessOrderRequest{}
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

	order := &entity.Order{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Type:      "receive",
	}

	orderID, err := h.orderUsecase.ProcessOrder(ctx.Request.Context(), order)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := dto.ProcessOrderResponse{
		OrderID: orderID,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "order processed",
		Data:    res,
	})
}

func (h *orderHandler) GetOrderByID(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderUsecase.GetOrderByID(ctx.Request.Context(), orderID)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := dto.OrderItem{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Type:      order.Type,
		CreatedAt: order.CreatedAt,
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "get order success",
		Data:    res,
	})
}

func (h *orderHandler) GetOrders(ctx *gin.Context) {
	req := dto.GetOrdersRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
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
	req.DefaultIfEmpty()

	paginatorParam := &entity.PaginationParam{
		OrderBy:        req.SortedBy,
		OrderDirection: req.Sort,
		Limit:          req.Limit,
		Offset:         (req.Page - 1) * req.Limit,
	}

	orders, count, err := h.orderUsecase.GetOrders(ctx.Request.Context(), paginatorParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := []dto.OrderItem{}
	for _, o := range orders {
		item := dto.OrderItem{
			ID:        o.ID,
			ProductID: o.ProductID,
			Quantity:  o.Quantity,
			Type:      o.Type,
			CreatedAt: o.CreatedAt,
		}
		res = append(res, item)
	}

	paginator := dto.MappingPaginator(req.Page, req.Limit, count)
	ctx.JSON(http.StatusOK, dto.Response{
		Message:   "get orders success",
		Data:      res,
		Paginator: &paginator,
	})
}
