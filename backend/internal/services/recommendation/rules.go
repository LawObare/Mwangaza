// Package recommendation (rules.go) contains the top-level decision logic.
//
// It evaluates all environmental factors (soil moisture, temperature,
// rain probability, NDVI, wind speed) against defined thresholds
// and delegates to specialized sub-files (irrigation, rainfall, etc.).
//
// Thresholds (example):
//   Soil moisture < 15%   → irrigation needed
//   Rain probability > 60% → delay irrigation
//   Temperature > 35°C    → heat stress warning
//   NDVI < 0.3            → poor vegetation health
//
// Each check returns a Recommendation with a priority level and message.
package recommendation