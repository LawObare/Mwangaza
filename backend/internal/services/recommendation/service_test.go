package recommendation

import (
	"strings"
	"testing"

	"mwangaza/internal/models"
)

func TestGenerateRecommendationsRules(t *testing.T) {
	baseline := models.SatelliteData{SoilMoisture: 40, Temperature: 25, RainProbability: 10, NDVI: 0.7, WindSpeed: 5}
	tests := []struct {
		name       string
		data       models.SatelliteData
		typeName   string
		confidence string
	}{
		{"irrigation", models.SatelliteData{SoilMoisture: 19, Temperature: 25, RainProbability: 10, NDVI: 0.7, WindSpeed: 5}, "irrigation", "MEDIUM"},
		{"likely rain", models.SatelliteData{SoilMoisture: 40, Temperature: 25, RainProbability: 55, NDVI: 0.7, WindSpeed: 5}, "rainfall", "MEDIUM"},
		{"heavy rain", models.SatelliteData{SoilMoisture: 40, Temperature: 25, RainProbability: 80, NDVI: 0.7, WindSpeed: 5}, "rainfall", "HIGH"},
		{"temperature", models.SatelliteData{SoilMoisture: 40, Temperature: 33, RainProbability: 10, NDVI: 0.7, WindSpeed: 5}, "temperature", "MEDIUM"},
		{"ndvi", models.SatelliteData{SoilMoisture: 40, Temperature: 25, RainProbability: 10, NDVI: 0.34, WindSpeed: 5}, "ndvi", "MEDIUM"},
		{"wind", models.SatelliteData{SoilMoisture: 40, Temperature: 25, RainProbability: 10, NDVI: 0.7, WindSpeed: 31}, "wind", "HIGH"},
		{"drought stress", models.SatelliteData{SoilMoisture: 19, Temperature: 25, RainProbability: 10, NDVI: 0.34, WindSpeed: 5}, "drought_stress", "HIGH"},
		{"heat stress", models.SatelliteData{SoilMoisture: 19, Temperature: 35, RainProbability: 10, NDVI: 0.7, WindSpeed: 5}, "heat_stress", "HIGH"},
		{"monitoring", baseline, "monitoring", "LOW"},
	}

	farm := models.Farm{ID: 1, Crop: "Maize", PreferredLanguage: "English"}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recommendations := GenerateRecommendations(farm, test.data)
			for _, recommendation := range recommendations {
				if recommendation.Type == test.typeName {
					if recommendation.Confidence != test.confidence {
						t.Fatalf("confidence = %q, want %q", recommendation.Confidence, test.confidence)
					}
					return
				}
			}
			t.Fatalf("did not generate %q: %#v", test.typeName, recommendations)
		})
	}
}

func TestGenerateRecommendationsMultipleAndDeterministic(t *testing.T) {
	farm := models.Farm{ID: 2, Crop: "Tomatoes", PreferredLanguage: "English"}
	data := models.SatelliteData{SoilMoisture: 19, Temperature: 35, RainProbability: 80, NDVI: 0.34, WindSpeed: 31}
	recommendations := GenerateRecommendations(farm, data)
	want := []string{"drought_stress", "heat_stress", "rainfall", "irrigation", "temperature", "ndvi", "wind"}
	if len(recommendations) != len(want) {
		t.Fatalf("generated %d recommendations, want %d", len(recommendations), len(want))
	}
	for i, typeName := range want {
		if recommendations[i].Type != typeName {
			t.Fatalf("recommendation %d = %q, want %q", i, recommendations[i].Type, typeName)
		}
		if recommendations[i].CreatedAt != recommendations[0].CreatedAt {
			t.Fatalf("recommendations must share one generation timestamp")
		}
	}
}

func TestGenerateRecommendationsLanguageAndCropFallback(t *testing.T) {
	data := models.SatelliteData{SoilMoisture: 19, Temperature: 25, RainProbability: 10, NDVI: 0.7, WindSpeed: 5}
	english := GenerateRecommendations(models.Farm{Crop: "Tomatoes", PreferredLanguage: "English"}, data)[0]
	swahili := GenerateRecommendations(models.Farm{Crop: "Tomatoes", PreferredLanguage: "Swahili"}, data)[0]
	fallback := GenerateRecommendations(models.Farm{Crop: "Sorghum", PreferredLanguage: "English"}, data)[0]
	if !strings.Contains(english.Message, "tomato") {
		t.Fatalf("English crop-specific message missing: %q", english.Message)
	}
	if !strings.Contains(swahili.Message, "nyanya") {
		t.Fatalf("Swahili message missing: %q", swahili.Message)
	}
	if fallback.Message == "" || fallback.Message == english.Message {
		t.Fatalf("fallback crop message should be safe and non-empty: %q", fallback.Message)
	}
}

func TestHighestPriorityPreservesRuleOrderOnTie(t *testing.T) {
	first := models.Recommendation{Type: "first", Priority: "HIGH"}
	best := HighestPriority([]models.Recommendation{first, {Type: "second", Priority: "HIGH"}})
	if best.Type != "first" {
		t.Fatalf("highest priority tie = %q, want first", best.Type)
	}
}
