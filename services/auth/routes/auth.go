// routes/auth.go
package routes

import (
	"github.com/abeselom-personal/personal-finance/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, authController *controller.AuthController) {
	authGroup := router.Group("/auth")
	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// @Summary Register a new user
	// @Description Create a new user with email or phone and password
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User registration info"
	// @Success 201 {object} models.User
	// @Failure 400 {object} gin.H
	// @Router /auth/register [post]
	authGroup.POST("/register", authController.Register)

	// @Summary User login
	// @Description Authenticate user by email/phone and password
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param credentials body map[string]string true "Login credentials"
	// @Success 200 {object} map[string]string "Tokens"
	// @Failure 401 {object} gin.H
	// @Router /auth/login [post]
	authGroup.POST("/login", authController.Login) // TODO

	// @Summary Logout user
	// @Description Revoke user token
	// @Tags Auth
	// @Produce json
	// @Success 200 {object} gin.H
	// @Failure 401 {object} gin.H
	// @Router /auth/logout [post]
	authGroup.POST("/logout", nil) // TODO
}
