package handlers

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"mwangaza/internal/config"
	"mwangaza/internal/middleware"
	"mwangaza/internal/models"
	"mwangaza/internal/services/auth"
	"mwangaza/internal/utils"
)

const (
	// tokenTTL is how long an issued JWT stays valid.
	tokenTTL = 24 * time.Hour
	// minPasswordLength matches the Flutter sign-up form validation.
	minPasswordLength = 6
)

// AuthHandler registers and authenticates accounts against the local store.
type AuthHandler struct {
	Service *auth.Service
	Config  config.Config
}

func NewAuthHandler(service *auth.Service, cfg config.Config) *AuthHandler {
	return &AuthHandler{Service: service, Config: cfg}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Register creates an account and returns a token, so the Flutter client can go
// straight to the dashboard after signing up.
//
// @Summary Register a new user
// @Description Creates a new user account and returns a JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param body body registerRequest true "Registration payload"
// @Success 201 {object} authResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := utils.DecodeJSON(r, &req); err != nil && !errors.Is(err, io.EOF) {
		utils.Error(w, http.StatusBadRequest, "invalid registration payload")
		return
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	if !looksLikeEmail(email) {
		utils.Error(w, http.StatusBadRequest, "a valid email address is required")
		return
	}
	if len(req.Password) < minPasswordLength {
		utils.Error(w, http.StatusBadRequest, "password must be at least 6 characters")
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		name = emailName(email)
	}

	user, err := h.Service.Register(name, email, strings.TrimSpace(req.Phone), req.Password)
	if err != nil {
		if errors.Is(err, auth.ErrEmailExists) {
			utils.Error(w, http.StatusConflict, "email already registered")
			return
		}
		utils.Error(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	response, err := h.issueToken(*user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	utils.Created(w, response)
}

// Login verifies credentials and returns a fresh token.
//
// @Summary Login
// @Description Authenticates user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body loginRequest true "Login credentials"
// @Success 200 {object} authResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := utils.DecodeJSON(r, &req); err != nil && !errors.Is(err, io.EOF) {
		utils.Error(w, http.StatusBadRequest, "invalid login payload")
		return
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	if email == "" || req.Password == "" {
		utils.Error(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, err := h.Service.Repository().GetUserByEmail(email)
	if err != nil || user == nil {
		utils.Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// GetUserByEmail returns the bcrypt hash in the Password field; it is
	// excluded from every JSON response by the model's `json:"-"` tag.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	response, err := h.issueToken(*user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	utils.Success(w, response)
}

// Me returns the account behind the request's Bearer token. It is mounted
// behind middleware.AuthMiddleware.Require.
//
// @Summary Current user
// @Description Returns the account identified by the request's bearer token
// @Tags auth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Router /auth/me [get]
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "authentication required")
		return
	}

	user, err := h.Service.Repository().GetUserByID(userID)
	if err != nil || user == nil {
		utils.Error(w, http.StatusUnauthorized, "account no longer exists")
		return
	}

	user.Password = ""
	utils.Success(w, *user)
}

func (h *AuthHandler) issueToken(user models.User) (authResponse, error) {
	secret := h.Config.JWTSecret
	if secret == "" {
		secret = config.DevJWTSecret
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(tokenTTL).Unix(),
	})

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return authResponse{}, err
	}

	user.Password = ""
	return authResponse{Token: signed, User: user}, nil
}

func looksLikeEmail(email string) bool {
	at := strings.Index(email, "@")
	if at <= 0 || at == len(email)-1 {
		return false
	}
	return strings.Contains(email[at+1:], ".") && !strings.ContainsAny(email, " \t")
}

func emailName(email string) string {
	if at := strings.Index(email, "@"); at > 0 {
		return email[:at]
	}
	return email
}
