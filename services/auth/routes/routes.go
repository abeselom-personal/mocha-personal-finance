// routes/routes.go
package routes

import (
	"github.com/abeselom-personal/personal-finance/controller"
	"github.com/abeselom-personal/personal-finance/db"
	"github.com/abeselom-personal/personal-finance/repositories"
	"github.com/abeselom-personal/personal-finance/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	//dependencies

	//repositories
	roleRepo := repositories.NewRoleRepository(db.DB)
	tokenRepo := repositories.NewTokenRepository(db.DB)
	userRepo := repositories.NewUserRepository(db.DB)

	//services
	authService := service.NewAuthService(userRepo, roleRepo, tokenRepo)

	//controllers
	authController := controller.NewAuthController(authService)

	//routes
	RegisterAuthRoutes(router, authController)
	RegisterUserRoutes(router)
	RegisterRoleRoutes(router)
	RegisterPermissionRoutes(router)
	RegisterTokenRoutes(router)
}
