package router

import (
	"warehouse-management-system/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine, authHandler handler.AuthHandler) *gin.RouterGroup {
	
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	return nil
}
