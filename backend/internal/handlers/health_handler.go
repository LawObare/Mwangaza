package handlers

import "github.com/gin-gonic/gin"

// HealthCheck  godoc
// @Summary     Health check
// @Description Returns server status
// @Tags        system
// @Produce     json
// @Success     200  {object}  map[string]string
// @Router      /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, map[string]string{"status": "ok"})
}