package jwtutils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type CustomClaims struct {
	ClientID int    `json:"client_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserClaims struct {
	CustomClaims
	jwt.RegisteredClaims
}

var accessTokenSecretKey = viper.GetString("ACCESS_TOKEN_SECRET")

func GenerateJWT(client CustomClaims) (string, error) {

	expiryHour := viper.GetInt("ACCESS_TOKEN_EXPIRY_HOUR")
	jwtIssuer := viper.GetString("JWT_ISSUER")
	jwtAudience := viper.GetString("JWT_AUDIENCE")
	now := time.Now()
	claims := UserClaims{
		CustomClaims: CustomClaims{
			ClientID: client.ClientID,
			Email:    client.Email,
			Role:     client.Role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  jwtIssuer,
			Subject: fmt.Sprint(client.ClientID),
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(time.Duration(expiryHour) * time.Hour),
			},
			Audience: jwt.ClaimStrings{jwtAudience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(accessTokenSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(JWT string) (UserClaims, error) {
	var claims UserClaims
	_, err := jwt.ParseWithClaims(JWT, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(accessTokenSecretKey), nil
	})
	if err != nil {
		return UserClaims{}, err
	}
	return claims, nil
}
