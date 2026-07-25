package alerts

import (
	"context"
	"log"
	"strings"
	"time"

	"mwangaza/internal/config"
	"mwangaza/internal/database"
	"mwangaza/internal/models"
	"mwangaza/internal/services/recommendation"
	"mwangaza/internal/services/sms"
)

type Runner struct {
	Store    *database.Store
	Config   config.Config
	Logger   *log.Logger
	generate func(*database.Store, config.Config, models.Farm) ([]models.Recommendation, models.SatelliteData, error)
}

type Summary struct {
	FarmsChecked int
	BatchesSaved int
	SMSSent      int
	Skipped      int
	Errors       int
}

func NewRunner(store *database.Store, cfg config.Config, logger *log.Logger) Runner {
	return Runner{Store: store, Config: cfg, Logger: logger}
}

func (r Runner) Start(ctx context.Context) {
	if !r.Config.AutoAlertsEnabled {
		return
	}

	interval := r.Config.AutoAlertsInterval
	if interval <= 0 {
		interval = 15 * time.Minute
	}

	go func() {
		r.logf("auto alerts enabled; checking farms every %s", interval)
		r.logSummary(r.RunOnce(ctx))

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				r.logf("auto alerts stopped")
				return
			case <-ticker.C:
				r.logSummary(r.RunOnce(ctx))
			}
		}
	}()
}

func (r Runner) RunOnce(ctx context.Context) Summary {
	var summary Summary
	for _, farm := range r.Store.ListFarms() {
		select {
		case <-ctx.Done():
			return summary
		default:
		}

		summary.FarmsChecked++
		generate := r.generate
		if generate == nil {
			generate = recommendation.GenerateForFarm
		}
		recommendations, _, err := generate(r.Store, r.Config, farm)
		if err != nil {
			summary.Errors++
			r.logf("auto alert failed for farm %d: %v", farm.ID, err)
			continue
		}
		summary.BatchesSaved++

		best := recommendation.HighestPriority(recommendations)
		if !shouldSend(best) || strings.TrimSpace(farm.Phone) == "" {
			summary.Skipped++
			continue
		}
		if r.withinCooldown(farm.ID, time.Now().UTC()) {
			summary.Skipped++
			continue
		}

		if _, err := sms.Send(r.Store, r.Config, models.SmsMessage{
			FarmID:      farm.ID,
			PhoneNumber: farm.Phone,
			Message:     best.Message,
		}); err != nil {
			summary.Errors++
			r.logf("auto SMS failed for farm %d: %v", farm.ID, err)
			continue
		}
		summary.SMSSent++
	}
	return summary
}

func shouldSend(recommendation models.Recommendation) bool {
	priority := strings.ToUpper(strings.TrimSpace(recommendation.Priority))
	return priority == "HIGH" || priority == "MEDIUM"
}

func (r Runner) withinCooldown(farmID int, now time.Time) bool {
	cooldown := r.Config.AutoAlertsSMSCooldown
	if cooldown <= 0 {
		return false
	}

	latest, ok := r.Store.LatestSMSForFarm(farmID)
	if !ok || latest.SentAt == "" {
		return false
	}

	sentAt, err := time.Parse(time.RFC3339, latest.SentAt)
	if err != nil {
		return false
	}
	return now.Sub(sentAt) < cooldown
}

func (r Runner) logSummary(summary Summary) {
	if summary.FarmsChecked == 0 {
		return
	}
	r.logf(
		"auto alerts checked=%d batches=%d sms=%d skipped=%d errors=%d",
		summary.FarmsChecked,
		summary.BatchesSaved,
		summary.SMSSent,
		summary.Skipped,
		summary.Errors,
	)
}

func (r Runner) logf(format string, args ...any) {
	if r.Logger != nil {
		r.Logger.Printf(format, args...)
	}
}
