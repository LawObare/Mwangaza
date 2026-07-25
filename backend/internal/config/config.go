package config

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Port         string
	DatabasePath string
	UseMockData  bool
	// KijaniDataURL is the exact Kijani endpoint that returns farm
	// environmental data. Its resource path is supplied by the Kijani contract.
	KijaniDataURL         string
	KijaniAPIKey          string
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
		KijaniDataURL:         getEnv("KIJANI_DATA_URL", ""),
		KijaniAPIKey:          getEnv("KIJANI_API_KEY", ""),
		SMSAPIKey:             getEnv("SMS_API_KEY", ""),
		SMSUsername:           getEnv("SMS_USERNAME", "sandbox"),
		SMSSenderID:           getEnv("SMS_SENDER_ID", ""),
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
	d, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return d
}

func ResolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Clean(path)
}
