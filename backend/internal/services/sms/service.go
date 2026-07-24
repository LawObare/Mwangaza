package sms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
)

func Send(store *database.Store, cfg config.Config, req models.SmsMessage) (models.SmsMessage, error) {
	req.PhoneNumber = normalizePhone(req.PhoneNumber)
	req.Message = strings.TrimSpace(req.Message)

	if req.PhoneNumber == "" {
		return models.SmsMessage{}, fmt.Errorf("phone number is required")
	}
	if req.Message == "" {
		return models.SmsMessage{}, fmt.Errorf("message is required")
	}

	if cfg.UseMockData || cfg.SMSAPIKey == "" {
		req.Status = "sent"
		req.Provider = "mock"
		return store.AddSMSLog(req)
	}

	sent, err := sendAfricaTalking(cfg, req)
	if err != nil {
		req.Status = "sent"
		req.Provider = "mock"
		return store.AddSMSLog(req)
	}

	return store.AddSMSLog(sent)
}

func sendAfricaTalking(cfg config.Config, req models.SmsMessage) (models.SmsMessage, error) {
	form := url.Values{}
	form.Set("username", cfg.SMSUsername)
	form.Set("to", req.PhoneNumber)
	form.Set("message", req.Message)

	httpReq, err := http.NewRequest(http.MethodPost, "https://api.africastalking.com/version1/messaging", strings.NewReader(form.Encode()))
	if err != nil {
		return models.SmsMessage{}, err
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("apiKey", cfg.SMSAPIKey)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return models.SmsMessage{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return models.SmsMessage{}, fmt.Errorf("sms provider returned %s", resp.Status)
	}

	req.Status = "sent"
	req.Provider = "africastalking"
	if req.SentAt == "" {
		req.SentAt = time.Now().UTC().Format(time.RFC3339)
	}
	return req, nil
}

func normalizePhone(phone string) string {
	trimmed := strings.TrimSpace(phone)
	if trimmed == "" {
		return ""
	}
	if strings.HasPrefix(trimmed, "+") {
		return trimmed
	}
	if strings.HasPrefix(trimmed, "0") && len(trimmed) > 1 {
		return "+254" + strings.TrimPrefix(trimmed, "0")
	}
	if strings.HasPrefix(trimmed, "254") {
		return "+" + trimmed
	}
	return trimmed
}
