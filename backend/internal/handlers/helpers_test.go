package handlers

import (
	"testing"

	"mwangaza/internal/models"
)

func TestStatusFromSatelliteTreatsZeroAsUnavailable(t *testing.T) {
	// OpenWeatherMap reports neither soil moisture nor NDVI; a zero for
	// those must not push a farm into "urgent" on its own.
	partial := models.SatelliteData{Temperature: 30, RainProbability: 10}
	if got := statusFromSatellite(partial); got != "healthy" {
		t.Fatalf("partial snapshot status = %q, want healthy", got)
	}
	lowSoil := models.SatelliteData{SoilMoisture: 10, Temperature: 30}
	if got := statusFromSatellite(lowSoil); got != "urgent" {
		t.Fatalf("low soil moisture status = %q, want urgent", got)
	}
	lowNDVI := models.SatelliteData{NDVI: 0.2, Temperature: 30}
	if got := statusFromSatellite(lowNDVI); got != "urgent" {
		t.Fatalf("low NDVI status = %q, want urgent", got)
	}
	hot := models.SatelliteData{Temperature: 36}
	if got := statusFromSatellite(hot); got != "urgent" {
		t.Fatalf("high temperature status = %q, want urgent", got)
	}
}
