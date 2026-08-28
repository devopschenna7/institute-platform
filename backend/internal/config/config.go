package config

import (
	"fmt"
	"os"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Environment string
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() (Config, error) {
	cfg := Config{
		App: AppConfig{
			Environment: getEnv("APP_ENV", "development"),
		},

		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnv("SERVER_PORT", "8080"),
		},

		Database: DatabaseConfig{
			Driver:   getEnv("DB_DRIVER", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "institute"),
		},
	}

	return cfg, validate(cfg)
}

func validate(cfg Config) error {
	if cfg.Server.Port == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}

	if cfg.Database.Driver == "" {
		return fmt.Errorf("DB_DRIVER is required")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
