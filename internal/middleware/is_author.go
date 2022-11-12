package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "role not defined",
			})
			c.Abort()

			return
		}

		if role != "author" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "not authorized to access this resource",
			})
			c.Abort()

			return
		}

		c.Next()
	}
}
