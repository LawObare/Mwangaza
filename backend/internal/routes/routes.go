package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"mwangaza/internal/handlers"
	"mwangaza/internal/middleware"
	"mwangaza/internal/services/auth"
)

func Setup(db *sql.DB) *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORS())

	// Initialize auth service
	authRepo := auth.NewSQLiteRepository(db)
	handlers.InitAuthService(authRepo)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/health", handlers.HealthCheck)

	api := r.Group("/api")
	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)

	protected := api.Group("")
	protected.Use(middleware.Auth())
	{
		protected.GET("/farms", handlers.ListFarms)
		protected.GET("/farms/:id", handlers.GetFarm)
		protected.GET("/satellite", handlers.GetSatelliteData)
		protected.GET("/recommendation", handlers.GetRecommendations)
		protected.GET("/sms", handlers.GetSMSHistory)
		protected.POST("/sms/send", handlers.SendSMS)
	}

	return r
}