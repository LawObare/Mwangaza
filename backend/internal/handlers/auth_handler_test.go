package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"mwangaza/internal/config"
	"mwangaza/internal/middleware"
	"mwangaza/internal/models"
	"mwangaza/internal/services/auth"
)

const testJWTSecret = "test-secret"

func TestAuthHandlerRegisterAndLogin(t *testing.T) {
	handler := newTestAuthHandler(t)

	register := httptest.NewRecorder()
	handler.Register(register, httptest.NewRequest(
		http.MethodPost,
		"/api/auth/register",
		stringsReader(`{"email":"Amina@Example.com","password":"secret123","phone":"+254700000001"}`),
	))
	if register.Code != http.StatusCreated {
		t.Fatalf("register status = %d, body = %s", register.Code, register.Body.String())
	}

	registered := decodeAuthResponse(t, register)
	if registered.Token == "" {
		t.Fatal("register did not return a token")
	}
	if registered.User.Email != "amina@example.com" {
		t.Fatalf("registered email = %q, want normalised address", registered.User.Email)
	}
	if registered.User.Name != "amina" {
		t.Fatalf("registered name = %q, want name derived from email", registered.User.Name)
	}
	if strings.Contains(register.Body.String(), "$2") {
		t.Fatalf("password hash leaked in response: %s", register.Body.String())
	}

	t.Run("duplicate email is rejected", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Register(recorder, httptest.NewRequest(
			http.MethodPost,
			"/api/auth/register",
			stringsReader(`{"email":"AMINA@example.com","password":"secret123"}`),
		))
		if recorder.Code != http.StatusConflict {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("short password is rejected", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Register(recorder, httptest.NewRequest(
			http.MethodPost,
			"/api/auth/register",
			stringsReader(`{"email":"short@example.com","password":"123"}`),
		))
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("invalid email is rejected", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Register(recorder, httptest.NewRequest(
			http.MethodPost,
			"/api/auth/register",
			stringsReader(`{"email":"not-an-email","password":"secret123"}`),
		))
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("login returns a token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Login(recorder, httptest.NewRequest(
			http.MethodPost,
			"/api/auth/login",
			stringsReader(`{"email":"amina@example.com","password":"secret123"}`),
		))
		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
		if response := decodeAuthResponse(t, recorder); response.Token == "" {
			t.Fatal("login did not return a token")
		}
	})

	t.Run("login rejects a wrong password", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Login(recorder, httptest.NewRequest(
			http.MethodPost,
			"/api/auth/login",
			stringsReader(`{"email":"amina@example.com","password":"wrong-password"}`),
		))
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("login rejects an unknown account", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler.Login(recorder, httptest.NewRequest(
			http.MethodPost,
			"/api/auth/login",
			stringsReader(`{"email":"nobody@example.com","password":"secret123"}`),
		))
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})
}

func TestAuthMiddlewareProtectsMe(t *testing.T) {
	handler := newTestAuthHandler(t)
	authMiddleware := middleware.NewAuth(testJWTSecret)

	register := httptest.NewRecorder()
	handler.Register(register, httptest.NewRequest(
		http.MethodPost,
		"/api/auth/register",
		stringsReader(`{"email":"grace@example.com","password":"secret123"}`),
	))
	token := decodeAuthResponse(t, register).Token

	guarded := authMiddleware.Require(http.HandlerFunc(handler.Me))

	t.Run("without a token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		guarded.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/api/auth/me", nil))
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("with a tampered token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/auth/me", nil)
		request.Header.Set("Authorization", "Bearer "+token+"tampered")
		guarded.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("with the issued token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/auth/me", nil)
		request.Header.Set("Authorization", "Bearer "+token)
		guarded.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
		}

		var response struct {
			Data models.User `json:"data"`
		}
		if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}
		if response.Data.Email != "grace@example.com" {
			t.Fatalf("me returned %#v", response.Data)
		}
	})
}

func TestAuthMiddlewareOptionalAllowsAnonymousRequests(t *testing.T) {
	handler := newTestAuthHandler(t)
	authMiddleware := middleware.NewAuth(testJWTSecret)

	anonymous := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if userID, ok := middleware.UserIDFromContext(r.Context()); ok {
			t.Errorf("anonymous request carried user id %d", userID)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	recorder := httptest.NewRecorder()
	authMiddleware.Optional(anonymous).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/api/farms", nil))
	if recorder.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want anonymous passthrough", recorder.Code)
	}

	register := httptest.NewRecorder()
	handler.Register(register, httptest.NewRequest(
		http.MethodPost,
		"/api/auth/register",
		stringsReader(`{"email":"john@example.com","password":"secret123"}`),
	))
	token := decodeAuthResponse(t, register).Token

	authed := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.UserIDFromContext(r.Context())
		if !ok || userID != 1 {
			t.Errorf("context user id = %d, ok = %v", userID, ok)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	recorder = httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/farms", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	authMiddleware.Optional(authed).ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNoContent {
		t.Fatalf("status = %d with a valid token", recorder.Code)
	}
}

func newTestAuthHandler(t *testing.T) *AuthHandler {
	t.Helper()

	store := testStore(t)
	return NewAuthHandler(auth.NewService(auth.NewStoreRepository(store)), config.Config{
		JWTSecret:   testJWTSecret,
		UseMockData: true,
	})
}

func decodeAuthResponse(t *testing.T, recorder *httptest.ResponseRecorder) authResponse {
	t.Helper()

	var envelope struct {
		Data authResponse `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("decode %s: %v", recorder.Body.String(), err)
	}
	return envelope.Data
}
