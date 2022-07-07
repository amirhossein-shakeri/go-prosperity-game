package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)

type AuthClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	// Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWTToken(u *User) string {
	claims := &AuthClaims{
		ID:    u.ID.Hex(),
		Email: u.Email,
		Name:  u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "go-prosperity-game-server",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(GetJWTSecret()))
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func ValidateJWTToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", t.Header["alg"])
		}
		return []byte(GetJWTSecret()), nil
	})
}

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "DEFAULT_SECRET_WCHICH_ANYBODY_CAN_GUESS_PROBABLY"
	}
	return secret
}

/*
Dude! I love this mechanical keyboard! It's awesome!
*/
