// routes/permissions.go
package routes

import "github.com/gin-gonic/gin"

func RegisterPermissionRoutes(router *gin.Engine) {
	permGroup := router.Group("/permissions")

	// @Summary List all permissions
	// @Description Get all permissions
	// @Tags Permissions
	// @Produce json
	// @Success 200 {array} models.Permission
	// @Router /permissions [get]
	permGroup.GET("", nil) // TODO

	// @Summary Create new permission
	// @Description Add a permission
	// @Tags Permissions
	// @Accept json
	// @Produce json
	// @Param permission body models.Permission true "Permission info"
	// @Success 201 {object} models.Permission
	// @Failure 400 {object} gin.H
	// @Router /permissions [post]
	permGroup.POST("", nil) // TODO

	// @Summary Update permission by ID
	// @Description Modify a permission
	// @Tags Permissions
	// @Accept json
	// @Produce json
	// @Param id path uint true "Permission ID"
	// @Param permission body models.Permission true "Permission info"
	// @Success 200 {object} models.Permission
	// @Failure 400 {object} gin.H
	// @Failure 404 {object} gin.H
	// @Router /permissions/{id} [put]
	permGroup.PUT("/:id", nil) // TODO

	// @Summary Delete permission by ID
	// @Description Remove a permission
	// @Tags Permissions
	// @Produce json
	// @Param id path uint true "Permission ID"
	// @Success 204 "No Content"
	// @Failure 404 {object} gin.H
	// @Router /permissions/{id} [delete]
	permGroup.DELETE("/:id", nil) // TODO
}
