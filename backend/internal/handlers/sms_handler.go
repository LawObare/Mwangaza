package handlers

import (
	"errors"
	"io"
	"net/http"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/recommendation"
	"mwangaza/internal/services/sms"
	"mwangaza/internal/utils"
)

type SMSHandler struct {
	Store  *database.Store
	Config config.Config
}

func NewSMSHandler(store *database.Store, cfg config.Config) *SMSHandler {
	return &SMSHandler{Store: store, Config: cfg}
}

func (h *SMSHandler) List(w http.ResponseWriter, r *http.Request) {
	utils.Success(w, h.Store.ListSMSLogs())
}

func (h *SMSHandler) Send(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FarmID      int    `json:"farm_id"`
		PhoneNumber string `json:"phone_number"`
		Message     string `json:"message"`
	}

	if err := utils.DecodeJSON(r, &req); err != nil && !errors.Is(err, io.EOF) {
		utils.Error(w, http.StatusBadRequest, "invalid sms payload")
		return
	}

	var farm models.Farm
	var found bool
	if req.FarmID > 0 {
		farm, found = h.Store.GetFarm(req.FarmID)
		if !found {
			utils.Error(w, http.StatusNotFound, "farm not found")
			return
		}
	}

	phone := req.PhoneNumber
	if phone == "" && found {
		phone = farm.Phone
	}

	message := req.Message
	if message == "" && found {
		if batch := h.Store.LatestRecommendationBatchForFarm(farm.ID); len(batch) > 0 {
			message = batch[0].Message
		} else {
			recommendations, _, err := recommendation.GenerateForFarm(h.Store, h.Config, farm)
			if err == nil {
				message = recommendation.HighestPriority(recommendations).Message
			}
		}
	}

	if phone == "" {
		utils.Error(w, http.StatusBadRequest, "phone number is required")
		return
	}
	if message == "" {
		utils.Error(w, http.StatusBadRequest, "message is required")
		return
	}

	sent, err := sms.Send(h.Store, h.Config, models.SmsMessage{
		FarmID:      req.FarmID,
		PhoneNumber: phone,
		Message:     message,
	})
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Created(w, sent)
}
