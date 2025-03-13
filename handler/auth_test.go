package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"warehouse-management-system/app/router"
	"warehouse-management-system/dto"
	"warehouse-management-system/entity"
	"warehouse-management-system/handler"
	"warehouse-management-system/mocks"
	validatorutils "warehouse-management-system/utils/validator"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_UserLogin(t *testing.T) {

	validLoginRequest := dto.LoginRequest{
		Email:    "ibrahim@example.com",
		Password: "12345678",
	}
	invalidLoginRequest := dto.LoginRequest{
		Email: "ibrahim",
	}

	validLoginCredential := entity.EmailPassword{
		Email:    "ibrahim@example.com",
		Password: "12345678",
	}
	// invalidLoginCredential := entity.LoginCredential{
	// 	Email:    "ibrahim",
	// 	Password: "12345678",
	// }
	loginToken := entity.LoginToken{
		UserID:      1,
		AccessToken: "jwt",
	}
	t.Run("should return 200 when logging in with correct credential", func(t *testing.T) {
		// given
		validatorutils.SetupValidator()
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(rec)

		mockAuthUsecase := mocks.NewAuthUsecase(t)
		mockAuthUsecase.On("Login", context.Background(), &validLoginCredential).Return(&loginToken, nil)

		authHandler := handler.NewAuthHandler(mockAuthUsecase)
		router.SetupAuthRouter(r, authHandler)
		expectedRes, _ := json.Marshal(gin.H{
			"data": gin.H{
				"access_token": loginToken.AccessToken,
				"id":           loginToken.UserID,
			},
			"message": "login success",
		})

		body, _ := json.Marshal(validLoginRequest)
		// when
		ctx.Request, _ = http.NewRequest("POST", "/login", strings.NewReader(string(body)))
		r.HandleContext(ctx)

		// then
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(expectedRes), rec.Body.String())
	})

	t.Run("should return 400 when logging in with no request body", func(t *testing.T) {
		// given
		validatorutils.SetupValidator()
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(rec)
		mockAuthUsecase := mocks.NewAuthUsecase(t)
		authHandler := handler.NewAuthHandler(mockAuthUsecase)
		router.SetupAuthRouter(r, authHandler)
		expectedRes, _ := json.Marshal(gin.H{
			"error": "invalid input",
		})

		// when
		ctx.Request, _ = http.NewRequest("POST", "/login", nil)
		r.HandleContext(ctx)

		// then
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, string(expectedRes), rec.Body.String())
	})

	t.Run("should return 400 when logging in with invalid input", func(t *testing.T) {
		// given
		validatorutils.SetupValidator()
		gin.SetMode(gin.TestMode)
		rec := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(rec)

		mockAuthUsecase := mocks.NewAuthUsecase(t)

		authHandler := handler.NewAuthHandler(mockAuthUsecase)
		router.SetupAuthRouter(r, authHandler)
		body, _ := json.Marshal(invalidLoginRequest)
		expectedRes, _ := json.Marshal(gin.H{
			"error": []gin.H{
				{
					"field":   "email",
					"message": "email must be a valid email address",
				},
				{
					"field":   "password",
					"message": "password is a required field",
				},
			},
		})

		// when
		ctx.Request, _ = http.NewRequest("POST", "/login", strings.NewReader(string(body)))
		r.HandleContext(ctx)

		// then
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, string(expectedRes), rec.Body.String())
	})
}
