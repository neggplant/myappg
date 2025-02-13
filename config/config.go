package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MongoDB struct {
		URI      string `yaml:"uri"`
		Database string `yaml:"database"`
	} `yaml:"mongodb"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config

func InitConfig() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(configFile, &AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
