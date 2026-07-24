package handlers

import (
	"github.com/gin-gonic/gin"
	"mwangaza/internal/models"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Register     godoc
// @Summary     Register a new user
// @Description Creates a new admin account
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       body  body      registerRequest  true  "Registration payload"
// @Success     201   {object}  models.User
// @Failure     400   {object}  map[string]interface{}
// @Router      /auth/register [post]
func Register(c *gin.Context) {
	c.JSON(201, models.User{})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Login        godoc
// @Summary     Login
// @Description Authenticates user and returns a JWT token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       body  body      loginRequest  true  "Login credentials"
// @Success     200   {object}  loginResponse
// @Failure     400   {object}  map[string]interface{}
// @Failure     401   {object}  map[string]interface{}
// @Router      /auth/login [post]
func Login(c *gin.Context) {
	c.JSON(200, loginResponse{})
}