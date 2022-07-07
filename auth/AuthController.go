package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"user": FindUser(ctx.GetString("user_id"))})
}

func Login(ctx *gin.Context) {
	req := &LoginRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		// return err.Error()
		return
	}

	user, message := AuthenticateUser(req.Email, req.Password)
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": message})
		// return message
		return
	}

	token := GenerateJWTToken(user)
	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": message, "user": user})
	// return token
}

func Signup(ctx *gin.Context) {
	req := &SignupRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	}

	user, err := CreateUser(req.Email, req.Name, req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Signed up successfully", "user": user, "token": GenerateJWTToken(user)})
}
