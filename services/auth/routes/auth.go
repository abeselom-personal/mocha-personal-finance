package routes

import (
	"net/http"

	"github.com/abeselom-personal/personal-finance/config"
	"github.com/abeselom-personal/personal-finance/controller"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RegisterAuthRoutes(
	router *gin.Engine,
	authController *controller.AuthController,
	authMiddleware gin.HandlerFunc,
	cfg *config.Config,
) {
	authGroup := router.Group("/auth")

	// Rate limiting
	registerLimiter := rate.NewLimiter(rate.Every(cfg.RateLimitWindow), cfg.RateLimitRegister)
	loginLimiter := rate.NewLimiter(rate.Every(cfg.RateLimitWindow), cfg.RateLimitLogin)

	authGroup.POST("/register", rateLimitMiddleware(registerLimiter), authController.Register)
	authGroup.POST("/login", rateLimitMiddleware(loginLimiter), authController.Login)
	authGroup.POST("/refresh", authController.Refresh)
	authGroup.POST("/logout", authMiddleware, authController.Logout)
}

func rateLimitMiddleware(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
