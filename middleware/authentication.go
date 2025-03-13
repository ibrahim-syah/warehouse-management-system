package middleware

import (
	"strings"
	"warehouse-management-system/sentinel"
	jwtutils "warehouse-management-system/utils/jwt"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		err := sentinel.ErrUnauthorized
		ctx.Error(err)
		return
	}

	splittedAuth := strings.Split(authorizationHeader, " ")
	if len(splittedAuth) != 2 || splittedAuth[0] != "Bearer" {
		err := sentinel.ErrUnauthorized
		ctx.Error(err)
		return
	}

	claims, err := jwtutils.ParseJWT(splittedAuth[1])
	if err != nil {
		err := sentinel.ErrUnauthorized
		ctx.Error(err)
		return
	}

	ctx.Set("UserID", claims.UserID)
	ctx.Set("Email", claims.Email)
	ctx.Set("Role", claims.Role)
}
