package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Constants
}

type Constants struct {
	Environment string
	Port        int

	Database struct {
		Dialect  string
		Host     string
		Port     int
		Username string
		Password string
		Name     string
		Charset  string
	}
}

func initViper() (Constants, error) {

	if os.Getenv("ENVIRONMENT") == "PRO" {
		viper.AutomaticEnv()
	} else {
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			return Constants{}, err
		}

		if err = viper.ReadInConfig(); err != nil {
			log.Panicf("Error reading config file, %s", err)
		}
	}

	var constants Constants
	var err = viper.Unmarshal(&constants)

	log.Printf("Loaded Env Variables:\n %+v", constants)
	return constants, err
}

func GetConfig() (*Config, error) {
	config := Config{}

	constants, err := initViper()
	config.Constants = constants

	return &config, err
}
