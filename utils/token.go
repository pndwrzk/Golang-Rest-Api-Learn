package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
    UserId uint64 `json:"user_id"`
    jwt.StandardClaims
}

const (
    accessTokenExpireDuration  = time.Hour
    refreshTokenExpireDuration = 7 * 24 * time.Hour // 1 week
)

func GenerateToken(id uint)(string, int64 ,error){
timer := time.Now().Add(accessTokenExpireDuration).Unix()

claims := jwt.MapClaims{}
claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = timer

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return "", timer,err
	}

	return tokenString, timer,nil


}

func GenerateRefreshToken(id uint)(string, int64,error){
timer := time.Now().Add(refreshTokenExpireDuration).Unix()

claims := jwt.MapClaims{}
claims["authorized"] = true
claims["id"] = id
claims["exp"] = timer
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

	if err != nil {
		return "", timer,err
	}

	return tokenString, timer,err


}

func ValidateRefreshToken(refreshToken string)(uint64, error){
	 token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
    })

	 if err != nil {
        return 0, err
    }

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
            return 0, jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
        }
        return claims.UserId, nil
    }

	return 0, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
}