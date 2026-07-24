package main

import (
	"log"
	"net/http"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/routes"
)

func main() {
	cfg := config.Load()

	store, err := database.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("open store: %v", err)
	}

	if err := database.Migrate(store); err != nil {
		log.Fatalf("migrate store: %v", err)
	}

	if err := database.Seed(store); err != nil {
		log.Fatalf("seed store: %v", err)
	}

	handler := routes.Setup(store, cfg)
	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Mwangaza backend listening on http://localhost:%s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("start server: %v", err)
	}
}
