package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/rvvg/okx-oi-exporter/config"
	"github.com/rvvg/okx-oi-exporter/metrics"
)

type OpenInterest struct {
    InstId   string `json:"instId"`
    InstType string `json:"instType"`
    OiUsd    string `json:"oiUsd"`
}

type ApiResponse struct {
    Code string         `json:"code"`
    Data []OpenInterest `json:"data"`
}


func FetchOpenInterest(cfg *config.Config, metrics *metrics.Metrics) {
    retryClient := retryablehttp.NewClient()
    retryClient.RetryMax = 3

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    req, err := retryablehttp.NewRequestWithContext(ctx, "GET", cfg.OKXEndpoint, nil)
    if err != nil {
        log.Printf("Error creating request: %v", err)
        return
    }

    resp, err := retryClient.Do(req)
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
        metrics.OpenInterestMetric.WithLabelValues(data.InstId, data.InstType).Set(oiusd)
    }
}

func CheckExchangeEndpoint(endpoint string) error {
    retryClient := retryablehttp.NewClient()
    retryClient.RetryMax = 3

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    req, err := retryablehttp.NewRequestWithContext(ctx, "GET", endpoint, nil)
    if err != nil {
        return err
    }

    resp, err := retryClient.Do(req)
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