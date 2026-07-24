# SpaceIoTBox Satellite Service

## Environment Variables

```env
SPACEIOTBOX_API_KEY=
SPACEIOTBOX_BASE_URL=
USE_MOCK_DATA=true
```

## Usage

```go
data, err := satellite.FetchSatelliteData()
```

## Example Output

```json
{
  "soil_moisture": 18.5,
  "temperature": 31,
  "rain_probability": 20,
  "ndvi": 0.64,
  "wind_speed": 8
}
```