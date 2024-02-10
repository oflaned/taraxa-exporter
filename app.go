package main

import (
	"github.com/gowizzard/dotenv"
	"github.com/oflaned/taraxa-exporter/internal/config"
	"github.com/oflaned/taraxa-exporter/internal/utils"
)



func main() {
	err := dotenv.Import(".env")
	if err != nil {
		panic("Unable to find .env file")
	}

	utils.InitLogger()
	utils.Logger.Info("Server started")

	configPath := dotenv.String("CONFIG_PATH")
	cfg := config.GetConfig(configPath)

	utils.Logger.Info("Api:", cfg.FilePaths)
}