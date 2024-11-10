package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rvvg/okx-oi-exporter/api"
	"github.com/rvvg/okx-oi-exporter/config"
)

func main() {
    log.Println("Starting OKX Open Interest Exporter")

    if err := api.CheckExchangeEndpoint(config.OKXEndpoint); err != nil {
        log.Fatalf("Failed to connect to exchange endpoint: %v", err)
    }

    http.Handle("/metrics", promhttp.Handler())

    go func() {
        for {
            api.FetchOpenInterest()
            time.Sleep(5 * time.Second)
        }
    }()

    log.Println("Exporter is ready to serve metrics on :8080/metrics")

    log.Fatal(http.ListenAndServe(":8080", nil))
}