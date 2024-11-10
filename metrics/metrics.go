package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
    OpenInterestMetric = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "okx_open_interest",
            Help: "Open interest data from OKX",
        },
        []string{"instId", "instType"},
    )
)

func init() {
    prometheus.MustRegister(OpenInterestMetric)
}