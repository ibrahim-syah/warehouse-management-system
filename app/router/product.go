package router

import (
	"warehouse-management-system/handler"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProductRouter(r *gin.Engine, productHandler handler.ProductHandler) *gin.RouterGroup {
	productGroup := r.Group("/products")

	productGroup.POST("", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("admin"), middleware.ErrorMiddleware, productHandler.AddProduct)
	productGroup.GET("", middleware.AuthenticationMiddleware, middleware.ErrorMiddleware, productHandler.GetProducts)
	productGroup.GET("/:id", middleware.AuthenticationMiddleware, middleware.ErrorMiddleware, productHandler.GetProductByID)
	productGroup.PUT("/:id", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("admin"), middleware.ErrorMiddleware, productHandler.UpdateProduct)
	productGroup.DELETE("/:id", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("admin"), middleware.ErrorMiddleware, productHandler.DeleteProduct)

	return productGroup
}
