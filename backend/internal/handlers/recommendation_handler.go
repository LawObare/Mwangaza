package handlers

import (
	"net/http"
	"strconv"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/recommendation"
	"mwangaza/internal/utils"
)

type RecommendationHandler struct {
	Store  *database.Store
	Config config.Config
}

func NewRecommendationHandler(store *database.Store, cfg config.Config) *RecommendationHandler {
	return &RecommendationHandler{Store: store, Config: cfg}
}

func (h *RecommendationHandler) List(w http.ResponseWriter, r *http.Request) {
	farmID, ok, err := parseQueryFarmID(r.URL.Query().Get("farm_id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var filter *int
	if ok {
		filter = &farmID
	}

	utils.Success(w, h.Store.ListRecommendations(filter))
}

func (h *RecommendationHandler) Generate(w http.ResponseWriter, r *http.Request) {
	farmID, ok, err := parseQueryFarmID(r.URL.Query().Get("farm_id"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok {
		farm, found := h.Store.GetFarm(farmID)
		if !found {
			utils.Error(w, http.StatusNotFound, "farm not found")
			return
		}

		rec, _, err := recommendation.GenerateForFarmWithToken(h.Store, h.Config, farm, bearerToken(r))
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.Created(w, rec)
		return
	}

	farms := h.Store.ListFarms()
	if len(farms) == 0 {
		utils.Error(w, http.StatusNotFound, "no farms available")
		return
	}

	result := make(map[string][]models.Recommendation, len(farms))
	for _, farm := range farms {
		recs, _, err := recommendation.GenerateForFarmWithToken(h.Store, h.Config, farm, bearerToken(r))
		if err != nil {
			continue
		}
		result[strconv.Itoa(farm.ID)] = recs
	}

	utils.Created(w, result)
}
