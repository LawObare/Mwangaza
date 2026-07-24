package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/utils"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		utils.Info(c.Request.Method + " " + c.Request.URL.Path + " " + time.Since(start).String())
	}
}