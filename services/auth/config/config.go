package config

import "os"

type Config struct {
	ConnString string
}

func Load() Config {
	return Config{
		ConnString: getEnv("DATABASE_URL", "postgres://pqgotest:password@localhost:5432/pqgotest?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
