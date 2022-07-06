package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email" xml:"email"`
	Password string `json:"password" form:"password" xml:"password"`
}

type SignupRequest struct {
	Email    string `json:"email" form:"email" xml:"email"`
	Password string `json:"password" form:"password" xml:"password"`
	Name     string `json:"name" form:"name" xml:"name"`
}

func GetInfo(ctx *gin.Context) {
	//
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
	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": message})
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
