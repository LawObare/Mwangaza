package config

import "os"

func IsMockEnabled() bool {
	return os.Getenv("USE_MOCK_DATA") == "true"
}

func SpaceIoTBoxAPIKey() string {
	return os.Getenv("SPACEIOTBOX_API_KEY")
}

func SpaceIoTBoxBaseURL() string {
	return os.Getenv("SPACEIOTBOX_BASE_URL")
}