package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		fmt.Println(" COming in Auth Middleware")
		if c.GetHeader("X-AUTH-API") == "" {
			//		c.AbortWithStatus(401)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "NOT AUTHORISED",
			})
		}

		c.Next()
	}

}
