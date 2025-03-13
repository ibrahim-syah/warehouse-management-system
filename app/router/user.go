package router

import (
	"warehouse-management-system/handler"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine, userHandler handler.UserHandler) *gin.RouterGroup {
	userGroup := r.Group("/users")
	userGroup.GET("me", middleware.AuthenticationMiddleware, middleware.ErrorMiddleware, userHandler.GetCurrentUser)
	userGroup.GET("", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("admin"), middleware.ErrorMiddleware, userHandler.GetUsers)

	return nil
}
