package main

import (
	"fmt"
	"log"

	"github.com/gowizzard/dotenv"
	"github.com/oflaned/taraxa-exporter/internal/config"
)



func main() {
	err := dotenv.Import(".env")
	if err != nil {
		log.Panic("Unable to find .env file")
	}


	configPath := dotenv.String("CONFIG_PATH")
	cfg := config.GetConfig(configPath)


	
	fmt.Println("Apis:")
	for _, path := range cfg.FilePaths {
		fmt.Println(path)
	}
}