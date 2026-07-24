// Package models (recommendation.go) defines the Recommendation struct.
//
// Represents a row in the recommendations table. The recommendation engine
// (services/recommendation) generates these based on satellite data and farm status.
//
// Fields:
//   ID        int    — primary key
//   FarmID    int    — foreign key to farms.id
//   Message   string — human-readable advice (e.g., "Irrigate for 20 minutes")
//   Priority  string — "low", "medium", or "high"
//   Reason    string — why this recommendation was made
//   CreatedAt string — ISO 8601 timestamp
package models