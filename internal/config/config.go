package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func MustInit() *Config {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Invalid PORT: %v", err)
	}

	return &Config{
		Port: port,
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
