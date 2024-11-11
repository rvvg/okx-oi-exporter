package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
    OpenInterestMetric *prometheus.GaugeVec
}

func NewMetrics() *Metrics {
    metrics := &Metrics{
        OpenInterestMetric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "okx_open_interest",
                Help: "Open interest data from OKX",
            },
            []string{"instId", "instType"},
        ),
    }
    prometheus.MustRegister(metrics.OpenInterestMetric)
    return metrics
}