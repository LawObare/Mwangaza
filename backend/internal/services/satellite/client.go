// Package satellite (client.go) handles HTTP communication with SpaceIoTBox.
//
// Function:
//   func FetchLiveData() ([]byte, error)
//
// Steps:
//   1. Read SPACEIOTBOX_API_KEY and SPACEIOTBOX_BASE_URL from env
//   2. Create an HTTP GET request to the base URL
//   3. Add X-API-Key header and Accept: application/json
//   4. Execute with a 10-second timeout
//   5. Read and return the raw response body as []byte
//
// No JSON parsing here. Returns raw bytes to service.go which passes
// them to parser.go. This keeps HTTP concerns isolated.
package satellite