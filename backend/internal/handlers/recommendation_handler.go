package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/internal/models"
)

// GetRecommendations  godoc
// @Summary           Get recommendations
// @Description       Returns all recommendations, optionally filtered by farm_id
// @Tags              recommendations
// @Produce           json
// @Param             farm_id  query     int  false  "Filter by farm ID"
// @Security          BearerAuth
// @Success           200  {object}  []models.Recommendation
// @Failure           401  {object}  map[string]interface{}
// @Router            /recommendation [get]
func GetRecommendations(c *gin.Context) {
	c.JSON(200, []models.Recommendation{})
}