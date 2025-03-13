package router

import (
	"warehouse-management-system/handler"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProductRouter(r *gin.Engine, productHandler handler.ProductHandler) *gin.RouterGroup {
	productGroup := r.Group("/products")

	productGroup.Use(middleware.AuthenticationMiddleware)

	return productGroup
}
