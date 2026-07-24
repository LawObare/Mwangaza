// Package models (farm.go) defines the Farm struct representing a row
// in the farms SQLite table.
//
// Fields:
//   ID        int     — primary key, autoincrement
//   Name      string  — farm name
//   Farmer    string  — farmer's full name
//   Phone     string  — farmer's phone number (+254 format)
//   Latitude  float64 — GPS latitude
//   Longitude float64 — GPS longitude
//
// Used by handlers to serialize JSON responses and by database queries
// to scan rows. JSON tags should match the Flutter model's expected keys.
package models