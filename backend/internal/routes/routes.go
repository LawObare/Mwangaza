package routes

import (
	"net/http"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/handlers"
	"mwangaza/internal/middleware"
)

func Setup(store *database.Store, cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	healthHandler := handlers.NewHealthHandler(cfg)
	farmHandler := handlers.NewFarmHandler(store, cfg)
	satelliteHandler := handlers.NewSatelliteHandler(store, cfg)
	recommendationHandler := handlers.NewRecommendationHandler(store, cfg)
	smsHandler := handlers.NewSMSHandler(store, cfg)
	authHandler := handlers.NewAuthHandler(store)

	// Auth routes (public, no JWT required)
	mux.HandleFunc("POST /api/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/auth/login", authHandler.Login)
	mux.HandleFunc("GET /api/auth/user", authHandler.GetUser)
	mux.HandleFunc("POST /api/auth/logout", authHandler.Logout)

	mux.HandleFunc("GET /api/health", healthHandler.Get)
	mux.HandleFunc("GET /api/farms", farmHandler.List)
	mux.HandleFunc("POST /api/farms", farmHandler.Create)
	mux.HandleFunc("GET /api/farms/{id}", farmHandler.Get)
	mux.HandleFunc("GET /api/satellite", satelliteHandler.Get)
	mux.HandleFunc("GET /api/recommendation", recommendationHandler.List)
	mux.HandleFunc("POST /api/recommendation", recommendationHandler.Generate)
	mux.HandleFunc("GET /api/sms", smsHandler.List)
	mux.HandleFunc("POST /api/sms/send", smsHandler.Send)

	return middleware.CORS(mux)
}
