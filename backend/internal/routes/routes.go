package routes

import (
	"net/http"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/handlers"
	"mwangaza/internal/middleware"
	"mwangaza/internal/services/auth"
)

func Setup(store *database.Store, cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	authService := auth.NewService(auth.NewStoreRepository(store))
	authMiddleware := middleware.NewAuth(cfg.JWTSecret)

	authHandler := handlers.NewAuthHandler(authService, cfg)
	healthHandler := handlers.NewHealthHandler(cfg)
	farmHandler := handlers.NewFarmHandler(store, cfg)
	satelliteHandler := handlers.NewSatelliteHandler(store, cfg)
	recommendationHandler := handlers.NewRecommendationHandler(store, cfg)
	smsHandler := handlers.NewSMSHandler(store, cfg)

	RegisterAuthRoutes(mux, authHandler, authMiddleware)

	mux.HandleFunc("GET /api/health", healthHandler.Get)
	mux.HandleFunc("GET /api/farms", farmHandler.List)
	mux.HandleFunc("POST /api/farms", farmHandler.Create)
	mux.HandleFunc("GET /api/farms/{id}", farmHandler.Get)
	mux.HandleFunc("GET /api/satellite", satelliteHandler.Get)
	mux.HandleFunc("GET /api/recommendation", recommendationHandler.List)
	mux.HandleFunc("POST /api/recommendation", recommendationHandler.Generate)
	mux.HandleFunc("GET /api/sms", smsHandler.List)
	mux.HandleFunc("POST /api/sms/send", smsHandler.Send)

	// Resource routes stay reachable without a token so the offline demo works;
	// a valid Bearer token is attached to the request context when supplied.
	return middleware.CORS(authMiddleware.Optional(mux))
}
