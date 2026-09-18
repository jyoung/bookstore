package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (c DBConfig) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Name,
	)
}

// Load loads configuration settings from environment variables
func Load() Config {
	return Config{
		Port: getEnv("PORT"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_USER"),
			Name:     getEnv("DB_NAME"),
		},
	}
}

func getEnv(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return "UNKNOWN"
}
