package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/internal/models"
)

// GetSMSHistory  godoc
// @Summary      Get SMS history
// @Description  Returns all sent SMS messages
// @Tags         sms
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  []models.SmsLog
// @Failure      401  {object}  map[string]interface{}
// @Router       /sms [get]
func GetSMSHistory(c *gin.Context) {
	c.JSON(200, []models.SmsLog{})
}

type smsRequest struct {
	FarmID      int    `json:"farm_id"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
}

// SendSMS      godoc
// @Summary     Send an SMS
// @Description Sends an advisory SMS to a farmer's phone
// @Tags        sms
// @Accept      json
// @Produce     json
// @Param       body  body      smsRequest  true  "SMS payload"
// @Security    BearerAuth
// @Success     201   {object}  map[string]interface{}
// @Failure     400   {object}  map[string]interface{}
// @Failure     401   {object}  map[string]interface{}
// @Router      /sms/send [post]
func SendSMS(c *gin.Context) {
	c.JSON(201, map[string]interface{}{})
}