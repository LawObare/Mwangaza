// Package satellite (parser.go) converts SpaceIoTBox JSON into models.SatelliteData.
//
// Function:
//   func ParseSatelliteData(data []byte) (models.SatelliteData, error)
//
// Steps:
//   1. Unmarshal []byte into the private apiResponse struct (defined in types.go)
//   2. Extract fields from the nested weather/soil/vegetation objects
//   3. Return a populated models.SatelliteData
//
// No HTTP calls. No env reading. Pure data transformation.
// The apiResponse struct maps to SpaceIoTBox's JSON shape:
//   { "weather": { "temperature", "rain_probability", "wind_speed" },
//     "soil": { "moisture" },
//     "vegetation": { "ndvi" } }
package satellite