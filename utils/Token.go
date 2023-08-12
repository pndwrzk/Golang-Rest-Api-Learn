package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
    accessTokenExpireDuration  = time.Hour
    refreshTokenExpireDuration = 7 * 24 * time.Hour // 1 week
    jwtSecret                  = "your-secret-key"
)

func GenerateToken(email string)(string, error){
timer := time.Now().Add(accessTokenExpireDuration).Unix()

claims := jwt.MapClaims{}
claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = timer

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, err


}

func GenerateRefreshToken(email string)(string, error){
timer := time.Now().Add(refreshTokenExpireDuration).Unix()

claims := jwt.MapClaims{}
claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = timer
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, err


}