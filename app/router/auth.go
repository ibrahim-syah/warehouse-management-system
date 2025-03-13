package router

import (
	"warehouse-management-system/handler"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine, authHandler handler.AuthHandler) *gin.RouterGroup {

	r.POST("/login", middleware.ErrorMiddleware, authHandler.Login)
	r.POST("/register", middleware.ErrorMiddleware, authHandler.Register)

	return nil
}
