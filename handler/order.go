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

type OrderHandler interface {
	ProcessShippingOrder(ctx *gin.Context)
	ProcessReceiveOrder(ctx *gin.Context)
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
