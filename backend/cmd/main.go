package main

import (
	_ "mwangaza/docs/swagger"
	"mwangaza/internal/config"
	"mwangaza/internal/routes"
	"mwangaza/internal/database"
)

// @title           Mwangaza API
// @version         1.0
// @description     Farm advisory dashboard API — satellite data, recommendations, SMS alerts.
// @contact.name    Mwangaza Team
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description               Enter "Bearer <token>"
func main() {
	cfg := config.Load()
	db, err := database.Connect(cfg.DatabasePath)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	r := routes.Setup(db)
	r.Run(":" + cfg.Port)
}