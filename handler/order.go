package handler

import (
	"encoding/json"
	"errors"
	"fmt"
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

	go func() {
		_, err := h.orderUsecase.ProcessOrder(ctx.Request.Context(), order)
		if err != nil {
			err = fmt.Errorf("Error processing order: %v", err)
			ctx.Error(err)
		}
	}()

	ctx.JSON(http.StatusAccepted, dto.Response{
		Message: "order is being processed",
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

	go func() {
		_, err := h.orderUsecase.ProcessOrder(ctx.Request.Context(), order)
		if err != nil {
			err = fmt.Errorf("Error processing order: %v", err)
			ctx.Error(err)
		}
	}()

	ctx.JSON(http.StatusAccepted, dto.Response{
		Message: "order is being processed",
	})
}
