// routes/tokens.go
package routes

import "github.com/gin-gonic/gin"

func RegisterTokenRoutes(router *gin.Engine) {
	tokenGroup := router.Group("/tokens")

	// @Summary List active tokens
	// @Description Get list of user's active tokens
	// @Tags Tokens
	// @Produce json
	// @Success 200 {array} models.Token
	// @Failure 401 {object} gin.H
	// @Router /tokens [get]
	tokenGroup.GET("", nil) // TODO

	// @Summary Revoke token by ID
	// @Description Revoke a specific token manually
	// @Tags Tokens
	// @Produce json
	// @Param id path uint true "Token ID"
	// @Success 200 {object} gin.H
	// @Failure 404 {object} gin.H
	// @Router /tokens/{id} [delete]
	tokenGroup.DELETE("/:id", nil) // TODO
}
