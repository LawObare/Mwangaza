package config

import (
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Port              string
	DatabasePath      string
	UseMockData       bool
	SpaceIoTBoxAPIKey string
	SpaceIoTBoxBaseURL string
	SMSAPIKey         string
	SMSUsername       string
	SMSSenderID       string
	DefaultLanguage   string
}

func Load() Config {
	_ = loadDotEnvFiles(".env", "backend/.env")

	return Config{
		Port:               getEnv("PORT", "8080"),
		DatabasePath:       getEnv("DATABASE_PATH", "./data/lakenet.db"),
		UseMockData:        parseBool(getEnv("USE_MOCK_DATA", "true")),
		SpaceIoTBoxAPIKey:  getEnv("SPACEIOTBOX_API_KEY", ""),
		SpaceIoTBoxBaseURL: getEnv("SPACEIOTBOX_BASE_URL", ""),
		SMSAPIKey:          getEnv("SMS_API_KEY", ""),
		SMSUsername:        getEnv("SMS_USERNAME", "sandbox"),
		SMSSenderID:        getEnv("SMS_SENDER_ID", "Mwangaza"),
		DefaultLanguage:    getEnv("DEFAULT_LANGUAGE", "English"),
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

func ResolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Clean(path)
}
