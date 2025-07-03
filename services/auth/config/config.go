package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ConnString             string
	JWTSecret              string
	BcryptCost             int
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
	RateLimitRegister      int
	RateLimitLogin         int
	RateLimitWindow        time.Duration
	CookieDomain           string
	GIN_MODE               string
}

var Cfg Config

func LoadConfig() Config {
	bcryptCost, _ := strconv.Atoi(getEnv("BCRYPT_COST", "14"))
	accessExp, _ := strconv.Atoi(getEnv("ACCESS_TOKEN_EXP_MIN", "15"))
	refreshExp, _ := strconv.Atoi(getEnv("REFRESH_TOKEN_EXP_HOUR", "720")) // 30 days
	rateLimitReg, _ := strconv.Atoi(getEnv("RATE_LIMIT_REGISTER", "5"))
	rateLimitLogin, _ := strconv.Atoi(getEnv("RATE_LIMIT_LOGIN", "10"))
	rateWindow, _ := strconv.Atoi(getEnv("RATE_LIMIT_WINDOW_MIN", "1"))
	Cfg = Config{
		JWTSecret:              getEnv("JWT_SECRET", "default_secret"),
		BcryptCost:             bcryptCost,
		AccessTokenExpiration:  time.Duration(accessExp) * time.Minute,
		RefreshTokenExpiration: time.Duration(refreshExp) * time.Hour,
		RateLimitRegister:      rateLimitReg,
		RateLimitLogin:         rateLimitLogin,
		RateLimitWindow:        time.Duration(rateWindow) * time.Minute,
		ConnString:             getEnv("DATABASE_URL", "postgres://pqgotest:password@localhost:5432/pqgotest?sslmode=disable"),
		CookieDomain:           getEnv("COOKIE_DOMAIN", "localhost"),
		GIN_MODE:               getEnv("GIN_MODE", "debug"),
	}
	return Cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
