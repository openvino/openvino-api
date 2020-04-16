package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Constants
}

type Constants struct {
	Environment string
	Port        int
	Database Database
}

type Database struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func New() *Config {
	return &Config {
		Constants: Constants {
			Environment: getEnv("ENVIRONMENT", "DEV"),
			Port: getEnvInt("API_PORT", 7878),
			Database: Database {
				Dialect: getEnv("DATABASE_DIALECT", "mysql"),
				Host: getEnv("DATABASE_HOST", "database"),
				Port: getEnvInt("DATABASE_PORT", 3306),
				Username: getEnv("DATABASE_USERNAME", "test"),
				Password: getEnv("DATABASE_PASSWORD", "test123"),
				Name: getEnv("DATABASE_NAME", "test_db"),
				Charset: getEnv("DATABASE_CHARSET", "utf8"),
			},
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists { 
		return value 
	} else {
		log.Println("Default value used for " + key + ": " + defaultVal)
		return defaultVal
	}
}

func getEnvInt(key string, defaultVal int) int {
	valueStr := getEnv(key, strconv.Itoa(defaultVal));
	if value, err := strconv.Atoi(valueStr); err == nil { 
		return value 
	} else {
		return defaultVal
	}
}

