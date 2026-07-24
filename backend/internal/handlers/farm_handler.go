package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/internal/models"
)

// ListFarms     godoc
// @Summary      List all farms
// @Description  Returns every farm in the database
// @Tags         farms
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  []models.Farm
// @Failure      401  {object}  map[string]interface{}
// @Router       /farms [get]
func ListFarms(c *gin.Context) {
	c.JSON(200, []models.Farm{})
}

// GetFarm      godoc
// @Summary     Get a single farm
// @Description Returns one farm by its ID
// @Tags        farms
// @Produce     json
// @Param       id   path      int  true  "Farm ID"
// @Security    BearerAuth
// @Success     200  {object}  models.Farm
// @Failure     401  {object}  map[string]interface{}
// @Failure     404  {object}  map[string]interface{}
// @Router      /farms/{id} [get]
func GetFarm(c *gin.Context) {
	c.JSON(200, models.Farm{})
}