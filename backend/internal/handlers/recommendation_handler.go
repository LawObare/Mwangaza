package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/utils"
)

var mockRecommendations = []gin.H{
	{"farm_id": 1, "type": "irrigation", "message": "Irrigate for 20 minutes", "severity": "medium", "created_at": "2025-01-15T08:00:00Z"},
	{"farm_id": 2, "type": "rainfall", "message": "High rain probability — delay irrigation", "severity": "low", "created_at": "2025-01-15T08:00:00Z"},
	{"farm_id": 3, "type": "temperature", "message": "Heat stress detected — increase watering", "severity": "high", "created_at": "2025-01-15T08:00:00Z"},
}

func GetRecommendations(c *gin.Context) {
	utils.Success(c, mockRecommendations)
}