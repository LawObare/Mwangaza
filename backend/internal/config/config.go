package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	MockData    bool
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		MockData:    os.Getenv("USE_MOCK_DATA") == "true",
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}