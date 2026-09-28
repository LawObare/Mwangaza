# Satellite / Agro-Climate Service

Reads farm environmental data for the Mwangaza backend, either live from
OpenWeatherMap (or a legacy Kijani endpoint) or from a deterministic mock
when offline.

## Setup

Environment Variables (set in `backend/.env`):

```env
# Empty or USE_MOCK_DATA=true => mock data
WEATHER_DATA_URL=https://api.openweathermap.org/data/2.5/forecast
WEATHER_API_KEY=
USE_MOCK_DATA=true
```

## Usage

The rest of the backend only talks to this package through `FetchAndStore()`
and `FetchForFarm()`:

```go
import "mwangaza/internal/services/satellite"

snapshot, err := satellite.FetchAndStore(store, cfg, farm, bearerToken)
if err != nil {
    // Handle error
}
// Use snapshot (models.SatelliteData)
```

## Internal Workflow

1. **Mode selection** (`service.go`)
   * `USE_MOCK_DATA=true` or no `WEATHER_DATA_URL` => `MockForFarm()`.
   * Otherwise build the endpoint with the farm's `lat`/`lon`.
2. **Request**
   * OpenWeatherMap: the API key is appended as `appid` together with
     `units=metric`; no auth headers are sent, so the caller's JWT never
     leaves Mwangaza.
   * Legacy providers (Kijani): sends the caller's bearer token when
     present, otherwise the `X-API-Key` header.
   * Falls back to `MockForFarm()` on any network or HTTP error, so the demo
     never fails hard.
3. **Parse** (`service.go`)
   * OpenWeatherMap: aggregates the first 24 hours of the 3-hourly five-day
     forecast into the single risk snapshot the recommendation engine
     expects — max temperature (°C), max rain probability (`pop` as %), max
     wind speed (m/s converted to km/h). Soil moisture and NDVI are not
     reported and stay zero ("not available"); the rules skip metrics that
     were not reported.
   * Kijani: converts the five-day hourly agro-climate forecast the same way.
4. **Store**
   * Persists the snapshot against the farm in the file-backed store.

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
