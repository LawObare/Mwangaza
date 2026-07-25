package satellite

import (
	"os"
)

func GetAPIKey() string {
	return os.Getenv("SPACEIOTBOX_API_KEY")
}

func GetBaseURL() string {
	url := os.Getenv("SPACEIOTBOX_BASE_URL")
	if url == "" {
		return "https://api.spaceiotbox.com/v1"
	}
	return url
}

func UseMockData() bool {
	return os.Getenv("USE_MOCK_DATA") == "true"
}
