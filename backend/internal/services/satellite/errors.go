package satellite

import "errors"

var (
	ErrInvalidAPIKey   = errors.New("invalid API key")
	ErrTimeout         = errors.New("request timeout")
	ErrInvalidResponse = errors.New("invalid response from API")
	ErrNoConnection    = errors.New("could not connect to API")
	ErrParsing         = errors.New("error parsing response")
)
