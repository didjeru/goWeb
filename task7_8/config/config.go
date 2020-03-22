package config

import (
	"encoding/json"
	"log"
	"os"
)

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Base     string `json:"base"`
	Table    string `json:"table"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
	Host     string         `json:"host"`
	Port     string         `json:"port"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	if configFile != nil {
		defer func() {
			if err := configFile.Close(); err != nil {
				log.Println(err)
			}
		}()
	}
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Println(err)
	}
	return config
}
