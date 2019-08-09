package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	ConnectionString string `json:"connectionString"`
}

type Config struct {
	Database Database
}

func LoadConfigurationFromFile(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		panic(fmt.Sprintf("Could not load config file: %v", err.Error()))
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func LoadConfigurationFromEnv() Config {
	connectionString := os.Getenv("CONNECTION_STRING")

	if connectionString == "" {
		panic("Could not read connectionString from environment.")
	}

	return Config{
		Database: Database{
			ConnectionString: connectionString,
		},
	}
}
