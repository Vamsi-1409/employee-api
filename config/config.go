package config

import "os"

type Config struct {
    Port        string
    DBHost      string
    DBPort      string
    DBUser      string
    DBPassword  string
    DBName      string
    JWTSecret   string
}

func Load() Config {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return Config{
		Port: port,
	}
}
