package router

import (
	"warehouse-management-system/handler"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupLocationRouter(r *gin.Engine, locationHandler handler.LocationHandler) *gin.RouterGroup {
	locationGroup := r.Group("/locations")

	locationGroup.POST("", middleware.AuthenticationMiddleware, middleware.AuthorizerMiddlewareGenerator("admin"), middleware.ErrorMiddleware, locationHandler.AddLocation)
	locationGroup.GET("", middleware.AuthenticationMiddleware, middleware.ErrorMiddleware, locationHandler.GetLocations)

	return locationGroup
}
