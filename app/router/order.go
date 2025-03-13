package router

import (
	"warehouse-management-system/handler"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupOrderRouter(r *gin.Engine, orderHandler handler.OrderHandler) *gin.RouterGroup {
	orderGroup := r.Group("/orders")

	orderGroup.POST("/receive", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("staff"), middleware.ErrorMiddleware, orderHandler.ProcessReceiveOrder)
	orderGroup.POST("/ship", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("staff"), middleware.ErrorMiddleware, orderHandler.ProcessShippingOrder)

	return orderGroup
}
