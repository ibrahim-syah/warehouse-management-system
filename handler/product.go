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

type ProductHandler interface {
	AddProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
}

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(
	productUsecase usecase.ProductUsecase,
) ProductHandler {
	handler := &productHandler{
		productUsecase: productUsecase,
	}

	return handler
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	req := dto.ProductRequest{}
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

	product := entity.Product{
		Name:       req.Name,
		SKU:        req.SKU,
		Quantity:   req.Quantity,
		LocationID: req.LocationID,
	}
	id, err := h.productUsecase.AddProduct(ctx.Request.Context(), &product)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: "product added successfully",
		Data:    id,
	})
}

func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	req := dto.ProductRequest{}
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

	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	product := entity.Product{
		ID:         productID,
		Name:       req.Name,
		SKU:        req.SKU,
		Quantity:   req.Quantity,
		LocationID: req.LocationID,
	}
	err = h.productUsecase.UpdateProduct(ctx.Request.Context(), &product)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "product updated successfully",
	})
}

func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.productUsecase.DeleteProduct(ctx.Request.Context(), productID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "product deleted successfully",
	})
}

func (h *productHandler) GetProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	product, err := h.productUsecase.GetProductByID(ctx.Request.Context(), productID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "get product success",
		Data:    product,
	})
}

func (h *productHandler) GetProducts(ctx *gin.Context) {
	req := dto.GetProductsRequest{}
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

	products, count, err := h.productUsecase.GetProducts(ctx.Request.Context(), paginatorParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := []dto.ProductItem{}
	for _, p := range products {
		item := dto.ProductItem{
			ID:         p.ID,
			Name:       p.Name,
			SKU:        p.SKU,
			Quantity:   p.Quantity,
			LocationID: p.LocationID,
			CreatedAt:  p.CreatedAt,
		}
		res = append(res, item)
	}

	paginator := dto.MappingPaginator(req.Page, req.Limit, count)
	ctx.JSON(http.StatusOK, dto.Response{
		Message:   "get products success",
		Data:      res,
		Paginator: &paginator,
	})
}
