package main

import (
	"fmt"
	"net/http"

	"github.com/gowizzard/dotenv"
	"github.com/oflaned/taraxa-exporter/internal/services"
	"github.com/oflaned/taraxa-exporter/internal/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)



func main() {
	err := dotenv.Import(".env")
	if err != nil {
		panic("Unable to find .env file")
	}

	utils.InitLogger()
	utils.Logger.Info("Server started")

	_, err = services.NodeMetrics(dotenv.String("METRICS_API"))
	if err != nil {
		fmt.Println(err)
	}

	prometheus.MustRegister(services.SyncStatus)
	prometheus.MustRegister(services.SyncTime)
	services.RecordMetrics()
	
	http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(dotenv.String("PORT"), nil)

}