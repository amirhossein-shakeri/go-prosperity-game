package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No authorization header provided"})
			return
		}
		authToken := authHeader[len(BEARER_SCHEMA):]
		token, err := ValidateJWTToken(authToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		if token.Valid {
			claims := token.Claims.(*AuthClaims)
			// claims := token.Claims.(jwt.MapClaims)
			ctx.Set("authenticated", true)
			ctx.Set("user_id", claims.ID)
			ctx.Set("jwt_claims", claims)
			// log.Println("CLAIMS:", claims) // maybe bind to request later
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
