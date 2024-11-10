package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
    OKXEndpoint   string
    ExporterPort  string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, checking environment variables")
	}

	OKXEndpoint = os.Getenv("OKX_ENDPOINT")
	ExporterPort = os.Getenv("EXPORTER_PORT")

	if OKXEndpoint == "" || ExporterPort == "" {
		log.Fatal("Required environment variables OKX_ENDPOINT and EXPORTER_PORT are not set in either .env file or environment")
	}
}
