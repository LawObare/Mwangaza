package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: extract Bearer token, validate JWT, set user context
		c.Next()
	}
}