// routes/user.go
package routes

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")

	// @Summary Get current user profile
	// @Description Get details of logged-in user
	// @Tags User
	// @Produce json
	// @Success 200 {object} models.User
	// @Failure 401 {object} gin.H
	// @Router /user/profile [get]
	userGroup.GET("/profile", nil) // TODO

	// @Summary Update current user profile
	// @Description Update logged-in user info
	// @Tags User
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User info to update"
	// @Success 200 {object} models.User
	// @Failure 400 {object} gin.H
	// @Failure 401 {object} gin.H
	// @Router /user/profile [put]
	userGroup.PUT("/profile", nil) // TODO
}
