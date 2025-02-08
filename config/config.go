package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	App struct {
		Port int `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		URI  string `yaml:"uri"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}

var AppConfig *Config

// LoadConfig loads configuration from the specified YAML file
func LoadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	AppConfig = &Config{}
	if err := decoder.Decode(AppConfig); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	log.Println("Configuration loaded successfully")
}
