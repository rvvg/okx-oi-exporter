package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rvvg/okx-oi-exporter/api"
	"github.com/rvvg/okx-oi-exporter/config"
	"github.com/rvvg/okx-oi-exporter/metrics"
)

func main() {
    cfg := config.LoadEnv()
    metrics := metrics.NewMetrics()
    
    log.Println("Starting OKX Open Interest Exporter")

    if err := api.CheckExchangeEndpoint(cfg.OKXEndpoint); err != nil {
        log.Fatalf("Failed to connect to exchange endpoint: %v", err)
    }

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        api.FetchOpenInterest(cfg, metrics)
        promhttp.Handler().ServeHTTP(w, r)
    })

    log.Printf("Exporter is ready to serve metrics on %s/metrics", cfg.ExporterPort)

    log.Fatal(http.ListenAndServe(cfg.ExporterPort, nil))
}