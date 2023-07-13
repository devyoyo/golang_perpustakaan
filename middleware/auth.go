package middleware

import (
	"net/http"
	"project_perpustakaan/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "request need access token",
			})

			c.Abort()
			return
		}

		// validate token
		nip, _, err := auth.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})

			c.Abort()
			return
		}

		c.Set("x-nip", nip)

		c.Next()

	}
}
