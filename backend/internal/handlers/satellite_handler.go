package handlers

import (
	"net/http"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/satellite"
	"mwangaza/internal/utils"
)

type SatelliteHandler struct {
	Store  *database.Store
	Config config.Config
}

func NewSatelliteHandler(store *database.Store, cfg config.Config) *SatelliteHandler {
	return &SatelliteHandler{Store: store, Config: cfg}
}

func (h *SatelliteHandler) Get(w http.ResponseWriter, r *http.Request) {
	if farmID, ok, err := parseQueryFarmID(r.URL.Query().Get("farm_id")); err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	} else if ok {
		farm, found := h.Store.GetFarm(farmID)
		if !found {
			utils.Error(w, http.StatusNotFound, "farm not found")
			return
		}
		data, err := satellite.FetchAndStore(h.Store, h.Config, farm)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.Success(w, data)
		return
	}

	if latest, found := h.Store.LatestSatellite(); found {
		utils.Success(w, latest)
		return
	}

	farms := h.Store.ListFarms()
	if len(farms) > 0 {
		data, err := satellite.FetchAndStore(h.Store, h.Config, farms[0])
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.Success(w, data)
		return
	}

	utils.Success(w, satellite.MockForFarm(models.Farm{Crop: "Maize"}))
}
