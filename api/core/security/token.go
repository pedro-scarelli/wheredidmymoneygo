package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(accountID string) string {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	secretKey := []byte(jwtSecret)

	claims := jwt.MapClaims{
		"sub": accountID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

