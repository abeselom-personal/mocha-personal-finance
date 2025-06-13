// routes/roles.go
package routes

import "github.com/gin-gonic/gin"

func RegisterRoleRoutes(router *gin.Engine) {
	roleGroup := router.Group("/roles")

	// @Summary List all roles
	// @Description Get list of all roles
	// @Tags Roles
	// @Produce json
	// @Success 200 {array} models.Role
	// @Router /roles [get]
	roleGroup.GET("", nil) // TODO

	// @Summary Create new role
	// @Description Create a new role
	// @Tags Roles
	// @Accept json
	// @Produce json
	// @Param role body models.Role true "Role info"
	// @Success 201 {object} models.Role
	// @Failure 400 {object} gin.H
	// @Router /roles [post]
	roleGroup.POST("", nil) // TODO

	// @Summary Update role by ID
	// @Description Update role details
	// @Tags Roles
	// @Accept json
	// @Produce json
	// @Param id path uint true "Role ID"
	// @Param role body models.Role true "Role info"
	// @Success 200 {object} models.Role
	// @Failure 400 {object} gin.H
	// @Failure 404 {object} gin.H
	// @Router /roles/{id} [put]
	roleGroup.PUT("/:id", nil) // TODO

	// @Summary Delete role by ID
	// @Description Remove a role
	// @Tags Roles
	// @Produce json
	// @Param id path uint true "Role ID"
	// @Success 204 "No Content"
	// @Failure 404 {object} gin.H
	// @Router /roles/{id} [delete]
	roleGroup.DELETE("/:id", nil) // TODO
}
