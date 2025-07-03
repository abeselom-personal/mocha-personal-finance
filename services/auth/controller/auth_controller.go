// controller/auth_controller.go
package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/abeselom-personal/personal-finance/config"
	"github.com/abeselom-personal/personal-finance/dto"
	"github.com/abeselom-personal/personal-finance/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

type AuthController struct {
	AuthService *service.AuthService
	Validator   *validator.Validate
	Config      *config.Config
}

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

func NewAuthController(
	authService *service.AuthService,
	validator *validator.Validate,
	cfg *config.Config,
) *AuthController {
	return &AuthController{
		AuthService: authService,
		Validator:   validator,
		Config:      cfg,
	}
}

// @Summary Register a new user
// @Description Create a new user account and return tokens (sets tokens in cookies)
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration info"
// @Success 201 {object} dto.TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := ac.Validator.Struct(req); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Details: errors,
		})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	tokens, err := ac.AuthService.Register(ctx, req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == service.ErrEmailExists || err == service.ErrUsernameExists {
			status = http.StatusConflict
		}
		c.JSON(status, ErrorResponse{Error: err.Error()})
		return
	}

	ac.setAuthCookies(c, tokens)
	c.JSON(http.StatusCreated, tokens)
}

// @Summary Login user
// @Description Authenticate user and return tokens (sets tokens in cookies)
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "User credentials"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := ac.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input format"})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	tokens, err := ac.AuthService.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials"})
		return
	}

	ac.setAuthCookies(c, tokens)
	c.JSON(http.StatusOK, tokens)
}

// @Summary Refresh access token
// @Description Get new access token using refresh token from cookie
// @Tags Auth
// @Produce json
// @Param refresh_token cookie string true "Refresh token"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} controller.ErrorResponse
// @Failure 401 {object} controller.ErrorResponse
// @Router /auth/refresh [post]
func (ac *AuthController) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Refresh token cookie missing"})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	tokens, err := ac.AuthService.RefreshToken(ctx, refreshToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid refresh token"})
		return
	}

	ac.setAuthCookies(c, tokens)
	c.JSON(http.StatusOK, tokens)
}

// @Summary Logout user
// @Description Revoke refresh token and clear cookies (reads refresh token from cookie)
// @Tags Auth
// @Produce json
// @Param refresh_token cookie string true "Refresh token to revoke"
// @Success 204
// @Failure 400 {object} controller.ErrorResponse
// @Failure 401 {object} controller.ErrorResponse
// @Router /auth/logout [post]
func (ac *AuthController) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Refresh token cookie missing"})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := ac.AuthService.RevokeToken(ctx, refreshToken); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid token"})
		return
	}

	ac.clearAuthCookies(c)
	c.Status(http.StatusNoContent)
}

func (ac *AuthController) setAuthCookies(c *gin.Context, tokens *dto.TokenResponse) {
	httpOnly := true
	domain := ac.Config.CookieDomain
	secure := ac.Config.GIN_MODE == "release"

	c.SetCookie("access_token", tokens.AccessToken, int(tokens.ExpiresIn), "/", domain, secure, httpOnly)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(7*24*time.Hour.Seconds()), "/", domain, secure, httpOnly)
	c.SetCookie("Authorization", "Bearer "+tokens.AccessToken, int(tokens.ExpiresIn), "/", domain, secure, httpOnly)
}

func (ac *AuthController) clearAuthCookies(c *gin.Context) {
	domain := ac.Config.CookieDomain
	c.SetCookie("access_token", "", -1, "/", domain, false, true)
	c.SetCookie("refresh_token", "", -1, "/", domain, false, true)
	c.SetCookie("Authorization", "", -1, "/", domain, false, true)
}
