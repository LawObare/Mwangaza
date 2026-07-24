package config

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Port                  string
	DatabasePath          string
	UseMockData           bool
	SpaceIoTBoxAPIKey     string
	SpaceIoTBoxBaseURL    string
	SMSAPIKey             string
	SMSUsername           string
	SMSSenderID           string
	DefaultLanguage       string
	AutoAlertsEnabled     bool
	AutoAlertsInterval    time.Duration
	AutoAlertsSMSCooldown time.Duration
}

func Load() Config {
	_ = loadDotEnvFiles(".env", "backend/.env")

	return Config{
		Port:                  getEnv("PORT", "8080"),
		DatabasePath:          getEnv("DATABASE_PATH", "./data/lakenet.db"),
		UseMockData:           parseBool(getEnv("USE_MOCK_DATA", "true")),
		SpaceIoTBoxAPIKey:     getEnv("SPACEIOTBOX_API_KEY", ""),
		SpaceIoTBoxBaseURL:    getEnv("SPACEIOTBOX_BASE_URL", ""),
		SMSAPIKey:             getEnv("SMS_API_KEY", ""),
		SMSUsername:           getEnv("SMS_USERNAME", "sandbox"),
		SMSSenderID:           getEnv("SMS_SENDER_ID", "Mwangaza"),
		DefaultLanguage:       getEnv("DEFAULT_LANGUAGE", "English"),
		AutoAlertsEnabled:     parseBool(getEnv("AUTO_ALERTS_ENABLED", "false")),
		AutoAlertsInterval:    parseDuration(getEnv("AUTO_ALERTS_INTERVAL", "15m"), 15*time.Minute),
		AutoAlertsSMSCooldown: parseDuration(getEnv("AUTO_ALERTS_SMS_COOLDOWN", "6h"), 6*time.Hour),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			return trimmed
		}
	}
	return fallback
}

func parseBool(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func parseDuration(value string, fallback time.Duration) time.Duration {
	duration, err := time.ParseDuration(strings.TrimSpace(value))
	if err != nil || duration <= 0 {
		return fallback
	}
	return duration
}

func ResolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Clean(path)
}
