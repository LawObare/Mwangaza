// Package satellite (types.go) defines structs for SpaceIoTBox API JSON decoding.
//
// These types are unexported (lowercase) — they only exist to deserialize
// the external API response. They are NOT used outside parser.go.
//
// Structs:
//   apiResponse — top-level: { "weather": {...}, "soil": {...}, "vegetation": {...} }
//   soil        — { "moisture": float64 }
//   weather     — { "temperature": float64, "rain_probability": float64, "wind_speed": float64 }
//   vegetation  — { "ndvi": float64 }
//
// models.SatelliteData is the public type used everywhere else.
// Do NOT duplicate SatelliteData fields here.
package satellite