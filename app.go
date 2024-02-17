package main

import (
	"fmt"

	"github.com/gowizzard/dotenv"
	"github.com/oflaned/taraxa-exporter/internal/services"
	"github.com/oflaned/taraxa-exporter/internal/utils"
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
}