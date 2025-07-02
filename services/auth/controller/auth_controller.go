package controller

import (
	"net/http"

	"github.com/abeselom-personal/personal-finance/dto"
	"github.com/abeselom-personal/personal-finance/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User registration info"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Router /auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := ac.AuthService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := ac.AuthService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
