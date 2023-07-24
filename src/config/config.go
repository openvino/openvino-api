package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config - Global config variables
var Config Constants

type EthereumConfig struct {
	InfuraSecretKey string
	Network         string
	ENS             string
}

// DatabaseConfig - Database configuration variables
type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

// Constants - Configuration variables structure
type Constants struct {
	Environment   string
	Port          string
	Secret        string
	Email         string
	EmailPassword string
	DashboardUrl string
	ServerUrl 	string
	EmailSmtp 	string
	EmailPort 	string
	Database      DatabaseConfig
	Ethereum      EthereumConfig
}

// New - Retrieve configuration based on environment variables
func New() Constants {
	godotenv.Load()

	return Constants{
		DashboardUrl:         getEnv("DASHBOARD_URL", ""),
		ServerUrl :  getEnv("SERVER_URL", ""),
		EmailSmtp : getEnv("EMAIL_SMPT", ""),
		EmailPort: getEnv("EMAIL_PORT", ""),
		Email:         getEnv("EMAIL", "example@example.com"),
		EmailPassword: getEnv("EMAIL_PASSWORD", "example4534"),
		Environment:   getEnv("ENVIRONMENT", "DEV"),
		Port:          getEnv("API_PORT", "3000"),
		Secret:        getEnv("SECRET_AUTH", "Secretillos"),
		Database: DatabaseConfig{
			Username:     getEnv("DB_USERNAME", "root"),
			Password:     getEnv("DB_PASSWORD", "root"),
			Host:         getEnv("DB_HOST", "127.0.0.1"),
			Port:         getEnv("DB_PORT", "3306"),
			DatabaseName: getEnv("DB_NAME", "enchainte"),
		},
		Ethereum: EthereumConfig{
			InfuraSecretKey: getEnv("ETH_INFURA_SECRET", ""),
			Network:         getEnv("ETH_NETWORK", ""),
			ENS:             getEnv("ETH_ENS", "rinkibino.eth"),
		},
	}
}

func getEnv(key string, defaultVal string) (value string) {
	if val, exists := os.LookupEnv(key); exists {
		value = val
	} else {
		value = defaultVal
	}

	log.Printf("  || %s => %s", key, value)

	return
}
