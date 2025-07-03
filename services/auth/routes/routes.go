// routes/routes.go
package routes

import (
	"github.com/abeselom-personal/personal-finance/config"
	"github.com/abeselom-personal/personal-finance/controller"
	"github.com/abeselom-personal/personal-finance/db"
	"github.com/abeselom-personal/personal-finance/dto"
	"github.com/abeselom-personal/personal-finance/middleware"
	"github.com/abeselom-personal/personal-finance/repositories"
	"github.com/abeselom-personal/personal-finance/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	//repositories
	cfg := &config.Cfg
	// Initialize repositories
	userRepo := *repositories.NewUserRepository(db.DB)
	tokenRepo := *repositories.NewTokenRepository(db.DB)
	roleRepo := *repositories.NewRoleRepository(db.DB)

	// Initialize services
	authService := service.NewAuthService(userRepo, tokenRepo, roleRepo, cfg)

	// Initialize validator
	validator := dto.NewValidator()

	// Initialize controllers
	authController := controller.NewAuthController(authService, validator, cfg)

	// Initialize middleware
	authMiddleware := middleware.AuthMiddleware(authService)

	//routes
	RegisterAuthRoutes(router, authController, authMiddleware, cfg)
	RegisterUserRoutes(router)
	RegisterRoleRoutes(router)
	RegisterPermissionRoutes(router)
	RegisterTokenRoutes(router)
}
