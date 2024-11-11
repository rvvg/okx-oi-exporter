package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    OKXEndpoint   string
    ExporterPort  string
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, checking environment variables")
	}

	config := &Config{
		OKXEndpoint:  os.Getenv("OKX_ENDPOINT"),
		ExporterPort: os.Getenv("EXPORTER_PORT"),
	}

	if config.OKXEndpoint == "" || config.ExporterPort == "" {
		log.Fatal("Required environment variables OKX_ENDPOINT and EXPORTER_PORT are not set in either .env file or environment")
	}

	return config
}
