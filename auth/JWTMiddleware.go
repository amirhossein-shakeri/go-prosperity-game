package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const BEARER_SCHEMA = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		authToken := authHeader[len(BEARER_SCHEMA):]
		token, err := ValidateJWTToken(authToken)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println(claims) // maybe bind to request later
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
