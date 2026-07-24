package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/utils"
)

var mockFarms = []gin.H{
	{"id": 1, "name": "Green Valley Farm", "farmer": "John Kamau", "latitude": -1.2921, "longitude": 36.8219, "status": "healthy", "soil_moisture": 18.5, "temperature": 31},
	{"id": 2, "name": "Sunrise Acres", "farmer": "Mary Wanjiku", "latitude": -0.3031, "longitude": 36.0800, "status": "needs attention", "soil_moisture": 12.3, "temperature": 33},
	{"id": 3, "name": "Hilltop Farm", "farmer": "Peter Omondi", "latitude": -1.2833, "longitude": 36.8167, "status": "urgent", "soil_moisture": 8.1, "temperature": 35},
}

func GetFarms(c *gin.Context) {
	utils.Success(c, mockFarms)
}

func GetFarm(c *gin.Context) {
	utils.Success(c, mockFarms[0])
}