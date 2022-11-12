package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func IsAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationField := c.Request.Header.Get("Authorization")
		if len(authorizationField) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "access token not found",
			})
			c.Abort()

			return
		}

		token := strings.Split(authorizationField, " ")[1]
		if !strings.Contains(token, ".") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid bearer token",
			})
			c.Abort()

			return
		}

		tokenParse, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
			if jwtToken.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unauthorized resource")
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()

			return
		}

		claims, ok := tokenParse.Claims.(jwt.MapClaims)
		if !ok || !tokenParse.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token. not authorized",
			})
			c.Abort()

			return
		}

		var tm time.Time
		switch exp := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(exp), 0)
		case json.Number:
			v, _ := exp.Int64()
			tm = time.Unix(v, 0)
		}

		if tm.Unix() < time.Now().Local().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token expired",
			})
			c.Abort()

			return
		}

		c.Set("role", claims["role"])
		c.Next()
	}
}
