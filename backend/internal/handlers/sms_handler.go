package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/backend/internal/utils"
)

var smsHistory = []gin.H{
	{"id": 1, "farm_id": 1, "phone_number": "+254712345678", "message": "Irrigate for 20 minutes today", "status": "sent", "sent_at": "2025-01-15T08:05:00Z"},
	{"id": 2, "farm_id": 2, "phone_number": "+254798765432", "message": "Delay irrigation — rain expected", "status": "sent", "sent_at": "2025-01-15T08:06:00Z"},
	{"id": 3, "farm_id": 3, "phone_number": "+25475551234", "message": "Heat stress alert — increase watering", "status": "failed", "sent_at": "2025-01-15T08:07:00Z"},
}

func GetSMSHistory(c *gin.Context) {
	utils.Success(c, smsHistory)
}

type sendSMSRequest struct {
	FarmID      int    `json:"farm_id" binding:"required"`
	Message     string `json:"message" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func SendSMS(c *gin.Context) {
	var req sendSMSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "invalid request body")
		return
	}
	if !utils.IsValidPhone(req.PhoneNumber) {
		utils.BadRequest(c, "invalid phone number format")
		return
	}
	if !utils.IsValidFarmID(req.FarmID) {
		utils.BadRequest(c, "invalid farm id")
		return
	}
	utils.Created(c, gin.H{"message": "SMS sent successfully", "farm_id": req.FarmID})
}