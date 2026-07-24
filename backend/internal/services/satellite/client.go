package satellite

import (
	"io"
	"net/http"
	"os"
	"time"
)

func FetchLiveData() ([]byte, error) {
	apiKey := os.Getenv("SPACEIOTBOX_API_KEY")
	baseURL := os.Getenv("SPACEIOTBOX_BASE_URL")

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}