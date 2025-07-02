package config

import "os"

type Config struct {
	ConnString string
	JWTString  string
}

var Cfg Config

func Load() Config {
	Cfg = Config{
		ConnString: getEnv("DATABASE_URL", "postgres://pqgotest:password@localhost:5432/pqgotest?sslmode=disable"),
		JWTString:  getEnv("JWTString", "secret"),
	}
	return Cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
