package handlers

import (
	"net/http"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/utils"
)

type HealthHandler struct {
	Config config.Config
}

func NewHealthHandler(cfg config.Config) *HealthHandler {
	return &HealthHandler{Config: cfg}
}

func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	utils.Success(w, map[string]any{
		"status":            "ok",
		"service":           "mwangaza",
		"mock_mode":         h.Config.UseMockData,
		"default_language":  h.Config.DefaultLanguage,
		"timestamp":         time.Now().UTC().Format(time.RFC3339),
	})
}
