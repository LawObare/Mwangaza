package handlers

import (
	"errors"
	"io"
	"net/http"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/recommendation"
	"mwangaza/internal/services/satellite"
	"mwangaza/internal/utils"
)

type FarmHandler struct {
	Store  *database.Store
	Config config.Config
}

func NewFarmHandler(store *database.Store, cfg config.Config) *FarmHandler {
	return &FarmHandler{Store: store, Config: cfg}
}

func (h *FarmHandler) List(w http.ResponseWriter, r *http.Request) {
	farms := h.Store.ListFarms()
	result := make([]models.Farm, 0, len(farms))
	for _, farm := range farms {
		result = append(result, enrichFarm(h.Store, farm))
	}
	utils.Success(w, result)
}

func (h *FarmHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, ok, err := parseQueryFarmID(r.PathValue("id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if !ok {
		utils.Error(w, http.StatusBadRequest, "farm id is required")
		return
	}

	farm, found := h.Store.GetFarm(id)
	if !found {
		utils.Error(w, http.StatusNotFound, "farm not found")
		return
	}

	utils.Success(w, enrichFarm(h.Store, farm))
}

func (h *FarmHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name              string  `json:"name"`
		Farmer            string  `json:"farmer"`
		Phone             string  `json:"phone"`
		County            string  `json:"county"`
		SubCounty         string  `json:"sub_county"`
		Crop              string  `json:"crop"`
		GrowthStage       string  `json:"growth_stage"`
		Latitude          float64 `json:"latitude"`
		Longitude         float64 `json:"longitude"`
		PreferredLanguage string  `json:"preferred_language"`
	}

	if err := utils.DecodeJSON(r, &req); err != nil && !errors.Is(err, io.EOF) {
		utils.Error(w, http.StatusBadRequest, "invalid farm payload")
		return
	}

	if req.Name == "" || req.Farmer == "" || req.Phone == "" || req.Crop == "" {
		utils.Error(w, http.StatusBadRequest, "name, farmer, phone, and crop are required")
		return
	}

	created, err := h.Store.CreateFarm(models.Farm{
		Name:              req.Name,
		Farmer:            req.Farmer,
		Phone:             req.Phone,
		County:            req.County,
		SubCounty:         req.SubCounty,
		Crop:              req.Crop,
		GrowthStage:       req.GrowthStage,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
		PreferredLanguage: req.PreferredLanguage,
	})
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	snapshot, err := satellite.FetchAndStore(h.Store, h.Config, created)
	if err == nil {
		rec := recommendation.GenerateRecommendation(created, snapshot)
		if _, recErr := h.Store.AddRecommendation(rec); recErr == nil {
			created.Status = statusFromSeverity(rec.Severity)
		}
		created.SoilMoisture = snapshot.SoilMoisture
		created.Temperature = snapshot.Temperature
		created.RainProbability = snapshot.RainProbability
		created.NDVI = snapshot.NDVI
		created.WindSpeed = snapshot.WindSpeed
	}

	utils.Created(w, enrichFarm(h.Store, created))
}
