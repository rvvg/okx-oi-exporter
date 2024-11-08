package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type OpenInterest struct {
    InstId  string `json:"instId"`
    InstType string `json:"instType"`
    Oi      string `json:"oi"`
    OiCcy   string `json:"oiCcy"`
    OiUsd   string `json:"oiUsd"`
    Ts      string `json:"ts"`
}

type ApiResponse struct {
    Code string         `json:"code"`
    Data []OpenInterest `json:"data"`
}

var (
    openInterestMetric = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "okx_open_interest",
            Help: "Open interest data from OKX",
        },
        []string{"instId", "instType"},
    )
)

var (
    okxOIEndpoint = "https://www.okx.com/api/v5/public/open-interest?instType=SWAP"
)

func init() {
    prometheus.MustRegister(openInterestMetric)
}

func fetchOpenInterest() {
    resp, err := http.Get(okxOIEndpoint)
    if err != nil {
        log.Printf("Error fetching data: %v", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading response body: %v", err)
        return
    }

    var apiResponse ApiResponse
    err = json.Unmarshal(body, &apiResponse)
    if err != nil {
        log.Printf("Error unmarshalling JSON: %v", err)
        return
    }

    for _, data := range apiResponse.Data {
        oiusd, err := strconv.ParseFloat(data.OiUsd, 64)
        if err != nil {
            log.Printf("Error parsing oi: %v", err)
            continue
        }
        openInterestMetric.WithLabelValues(data.InstId, data.InstType).Set(oiusd)
    }
}

func checkExchangeEndpoint() error {
    resp, err := http.Get(okxOIEndpoint)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
    }
    var apiResponse ApiResponse
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return fmt.Errorf("error decoding JSON response: %v", err)
    }
    return nil
}

func main() {
    log.Println("Starting OKX Open Interest Exporter")

    if err := checkExchangeEndpoint(); err != nil {
        log.Fatalf("Failed to connect to exchange endpoint: %v", err)
    }

    http.Handle("/metrics", promhttp.Handler())

    go func() {
        for {
            fetchOpenInterest()
            time.Sleep(5 * time.Second)
        }
    }()

    log.Println("Exporter is ready to serve metrics on :8080/metrics")

    log.Fatal(http.ListenAndServe(":8080", nil))
}