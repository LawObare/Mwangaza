package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"mwangaza/internal/database"
	"mwangaza/internal/utils"
)

// AuthHandler handles user registration, login, and profile endpoints
// using the standard net/http library (matching the rest of the backend).
type AuthHandler struct {
	Store *database.Store
}

func NewAuthHandler(store *database.Store) *AuthHandler {
	return &AuthHandler{Store: store}
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param body body object true "Registration payload"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 409 {object} utils.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if req.Email == "" || req.Password == "" {
		utils.Error(w, http.StatusBadRequest, "email and password are required")
		return
	}

	if len(req.Password) < 6 {
		utils.Error(w, http.StatusBadRequest, "password must be at least 6 characters")
		return
	}

	// Check if email already exists
	if _, found := h.Store.GetUserByEmail(req.Email); found {
		utils.Error(w, http.StatusConflict, "email already registered")
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "failed to process password")
		return
	}

	// If no name provided, derive from email
	name := req.Name
	if name == "" {
		name = req.Email
	}

	user, err := h.Store.CreateUser(name, req.Email, string(hash))
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	// Don't expose the password hash
	user.Password = ""

	utils.JSON(w, http.StatusCreated, utils.APIResponse{
		Success: true,
		Data: map[string]any{
			"token": token,
			"user":  user,
		},
	})
}

// Login godoc
// @Summary Login
// @Description Authenticates user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body object true "Login credentials"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if req.Email == "" || req.Password == "" {
		utils.Error(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, found := h.Store.GetUserByEmail(req.Email)
	if !found {
		utils.Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	// Don't expose the password hash
	user.Password = ""

	utils.Success(w, map[string]any{
		"token": token,
		"user":  user,
	})
}

// GetUser returns the authenticated user's profile.
func (h *AuthHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromToken(r)
	if userID == 0 {
		utils.Error(w, http.StatusUnauthorized, "invalid or missing token")
		return
	}

	user, found := h.Store.GetUserByID(userID)
	if !found {
		utils.Error(w, http.StatusNotFound, "user not found")
		return
	}

	user.Password = ""
	utils.Success(w, user)
}

// Logout is a no-op for stateless JWT but provided for client compatibility.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	utils.Success(w, map[string]string{"message": "logged out"})
}

func jwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-in-production"
	}
	return []byte(secret)
}

func generateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtSecret())
}

// userIDFromToken extracts the user ID from the Bearer JWT in the request.
func userIDFromToken(r *http.Request) int {
	tokenStr := bearerToken(r)
	if tokenStr == "" {
		return 0
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return jwtSecret(), nil
	})
	if err != nil || !token.Valid {
		return 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}

	if uid, ok := claims["user_id"].(float64); ok {
		return int(uid)
	}
	return 0
}
