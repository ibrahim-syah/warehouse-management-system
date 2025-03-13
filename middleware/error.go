package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"warehouse-management-system/dto"
	"warehouse-management-system/sentinel"
	validatorutils "warehouse-management-system/utils/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorMiddleware(ctx *gin.Context) {

	if len(ctx.Errors) > 0 {
		err := ctx.Errors.Last()
		if errors.Is(err, sentinel.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Error: "unauthorized access",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Error: "something went wrong, please contact the server administrator",
		})
		return
	}

	ctx.Next()

	if len(ctx.Errors) > 0 {
		err := ctx.Errors.Last()

		var ute *json.UnmarshalTypeError
		if errors.As(err, &ute) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: fmt.Sprintf("%s must be %s", ute.Field, ute.Type),
			})
			return
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ValidationErrorMsg, len(ve))
			for i, fe := range ve {
				trans := validatorutils.GetTranslator()
				out[i] = ValidationErrorMsg{Field: fe.Field(), Message: fe.Translate(trans)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: out,
			})
			return
		}

		if errors.Is(err, sentinel.ErrAlreadyExist) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: "item already exist",
			})
			return
		} else if errors.Is(err, sentinel.ErrInvalidInput) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: "invalid input",
			})
			return
		} else if errors.Is(err, sentinel.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: "item not found",
			})
			return
		} else if errors.Is(err, sentinel.ErrInvalidCredential) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: "invalid credential",
			})
			return
		} else if errors.Is(err, sentinel.ErrUsecaseError) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: err.Error(),
			})
			return
		} else if errors.Is(err, sentinel.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Error: "unauthorized access",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Error: "something went wrong, please contact the server administrator",
		})
		return
	}
}
