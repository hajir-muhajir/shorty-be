package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv  string
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func Load() *Config {
	return &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "127.0.0.1"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "user"),
		DBPass:  getEnv("DB_PASS", ""),
		DBName:  getEnv("DB_NAME", "shorty"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)
}

