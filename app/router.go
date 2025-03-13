package app

import (
	"warehouse-management-system/app/router"
	"warehouse-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, handlers *appHandlers) {
	r.Use(middleware.ErrorMiddleware)
	router.SetupAuthRouter(r, handlers.authHandler)
	router.SetupProductRouter(r, handlers.productHandler)
}
