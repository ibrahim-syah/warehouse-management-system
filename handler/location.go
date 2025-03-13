package handler

import (
	"net/http"
	"warehouse-management-system/dto"
	"warehouse-management-system/entity"
	"warehouse-management-system/usecase"

	"github.com/gin-gonic/gin"
)

type LocationHandler interface {
	AddLocation(ctx *gin.Context)
	GetLocations(ctx *gin.Context)
}

type locationHandler struct {
	locationUsecase usecase.LocationUsecase
}

func NewLocationHandler(
	locationUsecase usecase.LocationUsecase,
) LocationHandler {
	handler := &locationHandler{
		locationUsecase: locationUsecase,
	}

	return handler
}

func (h *locationHandler) AddLocation(ctx *gin.Context) {
	req := dto.LocationRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	location := entity.Location{
		Name:     req.Name,
		Capacity: req.Capacity,
	}
	id, err := h.locationUsecase.AddLocation(ctx.Request.Context(), &location)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: "location added successfully",
		Data:    id,
	})
}

func (h *locationHandler) GetLocations(ctx *gin.Context) {
	req := dto.GetLocationsRequest{}
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

	locations, count, err := h.locationUsecase.GetLocations(ctx.Request.Context(), paginatorParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := []dto.LocationItem{}
	for _, l := range locations {
		item := dto.LocationItem{
			ID:        l.ID,
			Name:      l.Name,
			Capacity:  l.Capacity,
			CreatedAt: l.CreatedAt,
		}
		res = append(res, item)
	}

	paginator := dto.MappingPaginator(req.Page, req.Limit, count)
	ctx.JSON(http.StatusOK, dto.Response{
		Message:   "get locations success",
		Data:      res,
		Paginator: &paginator,
	})
}
