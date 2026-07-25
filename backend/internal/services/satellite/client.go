package satellite

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// FetchLiveData communicates with the SpaceIoTBox API to retrieve satellite data.
func FetchLiveData() ([]byte, error) {
	apiKey := GetAPIKey()
	if apiKey == "" {
		return nil, ErrInvalidAPIKey
	}

	baseURL := GetBaseURL()
	
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Accept", "application/json")

	logSendingRequest()
	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrNoConnection
	}
	defer resp.Body.Close()

	logResponseReceived()

	switch resp.StatusCode {
	case http.StatusOK:
		// Continue
	case http.StatusBadRequest:
		return nil, ErrInvalidResponse
	case http.StatusUnauthorized, http.StatusForbidden:
		return nil, ErrInvalidAPIKey
	case http.StatusNotFound:
		return nil, ErrInvalidResponse
	case http.StatusInternalServerError:
		return nil, ErrInvalidResponse
	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrParsing
	}

	if len(body) == 0 {
		return nil, ErrInvalidResponse
	}

	return body, nil
}
