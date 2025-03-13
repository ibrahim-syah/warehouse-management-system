package middleware

import (
	"warehouse-management-system/sentinel"

	"github.com/gin-gonic/gin"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func AuthorizerMiddlewareGenerator(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("Role")

		if !stringInSlice(role, allowedRoles) {
			err := sentinel.ErrUnauthorized
			ctx.Error(err)
			return
		}

		ctx.Next()
	}
}
