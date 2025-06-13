// routes/routes.go
package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	RegisterAuthRoutes(router)
	RegisterUserRoutes(router)
	RegisterRoleRoutes(router)
	RegisterPermissionRoutes(router)
	RegisterTokenRoutes(router)
}
