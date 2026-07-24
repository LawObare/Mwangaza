package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/services/satellite"
	"mwangaza/backend/internal/utils"
)

func GetSatelliteData(c *gin.Context) {
	data, err := satellite.FetchSatelliteData()
	if err != nil {
		utils.InternalError(c, "failed to fetch satellite data")
		return
	}
	utils.Success(c, data)
}