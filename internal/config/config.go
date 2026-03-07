package config

import (
	"os"
	"time"
)

type Config struct {
	Port          string
	AllowedOrigin string
	SessionTTL    time.Duration
}

func Load() *Config {
	sessionTTL, err := time.ParseDuration(getEnv("SESSION_TTL", "1h"))
	if err != nil {
		panic(err)
	}
	return &Config{
		Port:          getEnv("PORT", "8080"),
		AllowedOrigin: getEnv("ALLOWED_ORIGINS", "*"),
		SessionTTL:    sessionTTL,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
