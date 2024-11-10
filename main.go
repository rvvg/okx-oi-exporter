package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rvvg/okx-oi-exporter/api"
	"github.com/rvvg/okx-oi-exporter/config"
)

func main() {
    log.Println("Starting OKX Open Interest Exporter")

    if err := api.CheckExchangeEndpoint(config.OKXEndpoint); err != nil {
        log.Fatalf("Failed to connect to exchange endpoint: %v", err)
    }

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        api.FetchOpenInterest()
        promhttp.Handler().ServeHTTP(w, r)
    })

    log.Println("Exporter is ready to serve metrics on :8080/metrics")

    log.Fatal(http.ListenAndServe(":8080", nil))
}