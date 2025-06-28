package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port       string
	MongoURI   string
	JWTSecret  string
	UploadPath string
}

func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	loadEnvFile()

	config := &Config{
		Port:       getEnv("PORT", "8080"),
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017/ctf_toolkit"),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		UploadPath: getEnv("UPLOAD_PATH", "./uploads"),
	}

	// Validate JWT secret
	if err := validateJWTSecret(config.JWTSecret); err != nil {
		return nil, err
	}

	return config, nil
}

func validateJWTSecret(secret string) error {
	if secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	if len(secret) < 32 {
		return fmt.Errorf("JWT_SECRET must be at least 32 characters long")
	}

	if secret == "your-super-secret-jwt-key-change-in-production" ||
		secret == "your-secret-key-change-in-production" {
		return fmt.Errorf("JWT_SECRET must be changed from default value")
	}

	return nil
}

func loadEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		return // .env file doesn't exist, that's okay
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				os.Setenv(key, value)
			}
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}