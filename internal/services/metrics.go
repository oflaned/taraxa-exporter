package services

import (
	"time"

	"github.com/gowizzard/dotenv"
	"github.com/oflaned/taraxa-exporter/internal/utils"
	"github.com/prometheus/client_golang/prometheus"
)

var (
    SyncStatus = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "SyncStatus",
        Help: "This metric represents a sync status where 1 is synced",
    })

    SyncTime = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "SyncTime",
        Help: "This metric represents a sync time",
    })
)

func RecordMetrics() {
    api := dotenv.String("METRICS_API")
    go func() {
	    for {
            metrics, err := NodeMetrics(api)
            if err != nil {
                utils.Logger.Error("Error while get metrics fromt api.", err)
            }

            if metrics.Synced {
                SyncStatus.Set(1);
            } else { 
                SyncStatus.Set(0)
            }
            SyncTime.Set(float64(metrics.SyncingSeconds))

            time.Sleep(1 * time.Second)
        }
    }()
}