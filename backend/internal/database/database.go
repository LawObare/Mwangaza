package database

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"mwangaza/internal/models"
)

type Store struct {
	mu                   sync.RWMutex
	path                 string
	Farms                []models.Farm           `json:"farms"`
	SatelliteData        []models.SatelliteData  `json:"satellite_data"`
	Recommendations      []models.Recommendation `json:"recommendations"`
	SMSLogs              []models.SmsMessage     `json:"sms_logs"`
	Users                []models.User           `json:"users"`
	nextFarmID           int
	nextSatelliteID      int
	nextRecommendationID int
	nextSMSID            int
	nextUserID           int
}

type dump struct {
	Farms           []models.Farm           `json:"farms"`
	SatelliteData   []models.SatelliteData  `json:"satellite_data"`
	Recommendations []models.Recommendation `json:"recommendations"`
	SMSLogs         []models.SmsMessage     `json:"sms_logs"`
	Users           []models.User           `json:"users"`
}

func Open(path string) (*Store, error) {
	if path == "" {
		path = "./data/lakenet.db"
	}

	store := &Store{path: path}
	if err := store.load(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) load() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}

	content, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.recomputeNextIDsLocked()
			return nil
		}
		return err
	}

	if len(content) == 0 {
		s.recomputeNextIDsLocked()
		return nil
	}

	var d dump
	if err := json.Unmarshal(content, &d); err != nil {
		s.recomputeNextIDsLocked()
		return nil
	}

	s.Farms = d.Farms
	s.SatelliteData = d.SatelliteData
	s.Recommendations = d.Recommendations
	s.SMSLogs = d.SMSLogs
	s.Users = d.Users
	s.recomputeNextIDsLocked()
	return nil
}

func (s *Store) recomputeNextIDsLocked() {
	s.nextFarmID = nextIDFromFarms(s.Farms)
	s.nextSatelliteID = nextIDFromSatelliteData(s.SatelliteData)
	s.nextRecommendationID = nextIDFromRecommendations(s.Recommendations)
	s.nextSMSID = nextIDFromSMSLogs(s.SMSLogs)
	s.nextUserID = nextIDFromUsers(s.Users)
}

func (s *Store) persistLocked() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}

	payload := dump{
		Farms:           s.Farms,
		SatelliteData:   s.SatelliteData,
		Recommendations: s.Recommendations,
		SMSLogs:         s.SMSLogs,
		Users:           s.Users,
	}

	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0o644)
}

func nextID[T any](items []T, picker func(T) int) int {
	maxID := 0
	for _, item := range items {
		if id := picker(item); id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

func nextIDFromFarms(items []models.Farm) int {
	return nextID(items, func(item models.Farm) int { return item.ID })
}

func nextIDFromSatelliteData(items []models.SatelliteData) int {
	return nextID(items, func(item models.SatelliteData) int { return item.ID })
}

func nextIDFromRecommendations(items []models.Recommendation) int {
	return nextID(items, func(item models.Recommendation) int { return item.ID })
}

func nextIDFromSMSLogs(items []models.SmsMessage) int {
	return nextID(items, func(item models.SmsMessage) int { return item.ID })
}

func (s *Store) ListFarms() []models.Farm {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]models.Farm, len(s.Farms))
	copy(out, s.Farms)
	sort.Slice(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})
	return out
}

func (s *Store) GetFarm(id int) (models.Farm, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, farm := range s.Farms {
		if farm.ID == id {
			return farm, true
		}
	}
	return models.Farm{}, false
}

func (s *Store) CreateFarm(farm models.Farm) (models.Farm, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if farm.ID == 0 {
		farm.ID = s.nextFarmID
		s.nextFarmID++
	}

	now := time.Now().UTC().Format(time.RFC3339)
	if farm.CreatedAt == "" {
		farm.CreatedAt = now
	}
	if farm.UpdatedAt == "" {
		farm.UpdatedAt = now
	}
	if farm.Status == "" {
		farm.Status = "new"
	}
	if farm.PreferredLanguage == "" {
		farm.PreferredLanguage = "English"
	}

	s.Farms = append(s.Farms, farm)
	if err := s.persistLocked(); err != nil {
		return models.Farm{}, err
	}
	return farm, nil
}

func (s *Store) LatestSatellite() (models.SatelliteData, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return latestSatelliteLocked(s.SatelliteData, 0)
}

func (s *Store) LatestSatelliteForFarm(farmID int) (models.SatelliteData, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return latestSatelliteLocked(s.SatelliteData, farmID)
}

func latestSatelliteLocked(items []models.SatelliteData, farmID int) (models.SatelliteData, bool) {
	var latest models.SatelliteData
	found := false

	for _, item := range items {
		if farmID > 0 && item.FarmID != farmID {
			continue
		}
		if !found || item.Timestamp > latest.Timestamp || (item.Timestamp == latest.Timestamp && item.ID > latest.ID) {
			latest = item
			found = true
		}
	}

	return latest, found
}

func (s *Store) AddSatelliteData(data models.SatelliteData) (models.SatelliteData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if data.ID == 0 {
		data.ID = s.nextSatelliteID
		s.nextSatelliteID++
	}
	if data.Timestamp == "" {
		data.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	s.SatelliteData = append(s.SatelliteData, data)
	if err := s.persistLocked(); err != nil {
		return models.SatelliteData{}, err
	}
	return data, nil
}

func (s *Store) ListRecommendations(farmID *int) []models.Recommendation {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]models.Recommendation, 0, len(s.Recommendations))
	for _, item := range s.Recommendations {
		if farmID != nil && item.FarmID != *farmID {
			continue
		}
		out = append(out, item)
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].CreatedAt == out[j].CreatedAt {
			return out[i].ID > out[j].ID
		}
		return out[i].CreatedAt > out[j].CreatedAt
	})
	return out
}

func (s *Store) LatestRecommendationForFarm(farmID int) (models.Recommendation, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var latest models.Recommendation
	found := false
	for _, item := range s.Recommendations {
		if item.FarmID != farmID {
			continue
		}
		if !found || item.CreatedAt > latest.CreatedAt || (item.CreatedAt == latest.CreatedAt && item.ID > latest.ID) {
			latest = item
			found = true
		}
	}
	return latest, found
}

func (s *Store) AddRecommendation(rec models.Recommendation) (models.Recommendation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if rec.ID == 0 {
		rec.ID = s.nextRecommendationID
		s.nextRecommendationID++
	}
	if rec.CreatedAt == "" {
		rec.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	}
	if rec.Severity == "" {
		rec.Severity = "low"
	}
	if rec.Priority == "" {
		rec.Priority = strings.ToUpper(rec.Severity)
	}

	s.Recommendations = append(s.Recommendations, rec)
	if err := s.persistLocked(); err != nil {
		return models.Recommendation{}, err
	}
	return rec, nil
}

func (s *Store) ListSMSLogs() []models.SmsMessage {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]models.SmsMessage, len(s.SMSLogs))
	copy(out, s.SMSLogs)
	sort.Slice(out, func(i, j int) bool {
		if out[i].SentAt == out[j].SentAt {
			return out[i].ID > out[j].ID
		}
		return out[i].SentAt > out[j].SentAt
	})
	return out
}

func (s *Store) LatestSMSForFarm(farmID int) (models.SmsMessage, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var latest models.SmsMessage
	found := false
	for _, item := range s.SMSLogs {
		if item.FarmID != farmID {
			continue
		}
		if !found || item.SentAt > latest.SentAt || (item.SentAt == latest.SentAt && item.ID > latest.ID) {
			latest = item
			found = true
		}
	}
	return latest, found
}

func (s *Store) AddSMSLog(log models.SmsMessage) (models.SmsMessage, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if log.ID == 0 {
		log.ID = s.nextSMSID
		s.nextSMSID++
	}
	if log.SentAt == "" {
		log.SentAt = time.Now().UTC().Format(time.RFC3339)
	}
	if log.Status == "" {
		log.Status = "sent"
	}
	if log.Provider == "" {
		log.Provider = "mock"
	}

	s.SMSLogs = append(s.SMSLogs, log)
	if err := s.persistLocked(); err != nil {
		return models.SmsMessage{}, err
	}
	return log, nil
}

// ── User helpers ────────────────────────────────────────────────────────

func nextIDFromUsers(items []models.User) int {
	return nextID(items, func(item models.User) int { return item.ID })
}

func (s *Store) CreateUser(name, email, passwordHash string) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := models.User{
		ID:           s.nextUserID,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}
	s.nextUserID++

	s.Users = append(s.Users, user)
	if err := s.persistLocked(); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Store) GetUserByEmail(email string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.Users {
		if strings.EqualFold(user.Email, email) {
			user.Password = user.PasswordHash // expose hash for bcrypt comparison
			return user, true
		}
	}
	return models.User{}, false
}

func (s *Store) GetUserByID(id int) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.Users {
		if user.ID == id {
			user.Password = user.PasswordHash // expose hash for bcrypt comparison
			return user, true
		}
	}
	return models.User{}, false
}
