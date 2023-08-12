package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
      
        authHeader  := c.GetHeader("Authorization")
        if authHeader  == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
            c.Abort()
            return
        }

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ivalid Token format"})
            c.Abort()
            return

		}
		  token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("TOKEN_SECRET")), nil
        })

		 if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

		 if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
            c.Abort()
            return
        } 
		 c.Next()
       
    }
}

