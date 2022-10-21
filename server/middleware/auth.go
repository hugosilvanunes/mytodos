package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hugosilvanunes/mytodos/config"
)

func Auth(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid authorization"})
			return
		}

		strs := strings.Split(authHeader, " ")

		if len(strs) != 2 || strs[0] != "Bearer" || strs[1] == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid authorization"})
			return
		}

		jwtReceived := strs[1]
		token, err := jwt.Parse(jwtReceived, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(cfg.ENV.JWTSecretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid authorization"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Set("userID", claims["jti"])
		c.Next()
	}
}
