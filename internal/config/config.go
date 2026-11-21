package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

// membaca satu env var dan mengembalikan error jika kosong
func required(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is required", key)
	}
	return value, nil
}

func Load() (*Config, error) {
	cfg := &Config{}
	var err error
	cfg.AppPort, err = required("APP_PORT")
	if err != nil {
		return nil, err
	}

	cfg.DBHost, err = required("DB_HOST")
	if err != nil {
		return nil, err
	}

	cfg.DBPort, err = required("DB_PORT")
	if err != nil {
		return nil, err
	}

	cfg.DBName, err = required("DB_NAME")
	if err != nil {
		return nil, err
	}

	cfg.DBUser, err = required("DB_USER")
	if err != nil {
		return nil, err
	}

	cfg.DBPassword, err = required("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
