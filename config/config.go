package config

import "os"

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	DBSSLMode  string
}

func Load() Config {
	port := getEnv("PORT", "8080")

	return Config{
		Port:       port,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),
		JWTSecret:  getEnv("JWT_SECRET", "my-secret-key"),
		DBSSLMode:  getEnv("DB_SSLMODE", "require"),
	}
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
