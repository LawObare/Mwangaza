# Satellite Integration Service

This package integrates the LAKENET backend with the SpaceIoTBox API.

## Setup

Environment Variables:

```env
SPACEIOTBOX_API_KEY=your_api_key_here
SPACEIOTBOX_BASE_URL=https://api.spaceiotbox.com/v1
USE_MOCK_DATA=true
```

## Usage

The rest of the backend should only interact with this service through `FetchSatelliteData()`:

```go
import "mwangaza/internal/services/satellite"

data, err := satellite.FetchSatelliteData()
if err != nil {
    // Handle error
}
// Use data (models.SatelliteData)
```

## Internal Workflow

1.  **Read Config**: Checks `USE_MOCK_DATA`.
2.  **Mode Selection**:
    *   If `true`: Returns static data from `mock.go`.
    *   If `false`:
        1.  Calls `FetchLiveData()` in `client.go`.
        2.  Adds authentication headers and executes the request.
        3.  Passes raw JSON to `ParseSatelliteData()` in `parser.go`.
        4.  Returns the converted `models.SatelliteData`.

## Example Output

```json
{
  "soil_moisture": 18.5,
  "temperature": 31.2,
  "rain_probability": 20,
  "ndvi": 0.63,
  "wind_speed": 7.4,
  "timestamp": "2026-07-24T12:00:00Z"
}
```

## Error Handling

Handles:
- Missing/Invalid API Key
- Network Errors & Timeouts
- Invalid JSON/Empty Responses
- HTTP Status Errors (400, 401, 403, 404, 500)
