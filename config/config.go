package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Port        int
	DatabaseURL string
	MaxFileSize int64
	UploadsDir  string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Invalid PORT value")
	}

	maxFileSize, err := strconv.ParseInt(getEnv("MAX_FILE_SIZE", "10485760"), 10, 64) // Default: 10 MB
	if err != nil {
		log.Fatal("Invalid MAX_FILE_SIZE value")
	}

	return &Config{
		Port:        port,
		DatabaseURL: getEnv("DATABASE_URL", ""),
		MaxFileSize: maxFileSize,
		UploadsDir:  getEnv("UPLOADS_DIR", "./uploads"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
