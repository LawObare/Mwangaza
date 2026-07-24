package routes

import (
	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/handlers"
	"mwangaza/backend/internal/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/farms", handlers.GetFarms)
		api.GET("/farms/:id", handlers.GetFarm)
		api.GET("/satellite", handlers.GetSatelliteData)
		api.GET("/recommendation", handlers.GetRecommendations)
		api.GET("/sms", handlers.GetSMSHistory)
		api.POST("/sms/send", handlers.SendSMS)
	}

	return r
}