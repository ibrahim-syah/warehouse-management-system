package app

import (
	"warehouse-management-system/app/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, handlers *appHandlers) {
	router.SetupProductRouter(r, handlers.productHandler)
}
