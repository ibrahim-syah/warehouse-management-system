package app

import (
	"warehouse-management-system/app/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, handlers *appHandlers) {

	router.SetupAuthRouter(r, handlers.authHandler)
	router.SetupUserRouter(r, handlers.userHandler)
	router.SetupProductRouter(r, handlers.productHandler)
	router.SetupLocationRouter(r, handlers.locationHandler)

}
