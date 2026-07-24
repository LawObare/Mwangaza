package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/internal/models"
)

// GetSatelliteData  godoc
// @Summary         Get satellite data
// @Description     Fetches latest environmental data from SpaceIoTBox (or mock)
// @Tags            satellite
// @Produce         json
// @Security        BearerAuth
// @Success         200  {object}  models.SatelliteData
// @Failure         401  {object}  map[string]interface{}
// @Failure         500  {object}  map[string]interface{}
// @Router          /satellite [get]
func GetSatelliteData(c *gin.Context) {
	c.JSON(200, models.SatelliteData{})
}