package main

import (
	"mwangaza/backend/internal/config"
	"mwangaza/backend/internal/routes"
	"mwangaza/backend/internal/utils"
)

func main() {
	cfg := config.Load()
	utils.Info("starting Mwangaza backend on port " + cfg.Port)

	r := routes.Setup()
	if err := r.Run(":" + cfg.Port); err != nil {
		utils.Fatal("server failed to start", err)
	}
}