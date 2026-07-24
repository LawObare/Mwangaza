package middleware

import (
	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/utils"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.Error("panic recovered", nil)
				c.AbortWithStatusJSON(500, gin.H{"success": false, "error": "internal server error"})
			}
		}()
		c.Next()
	}
}