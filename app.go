package main

import (
	"fmt"

	"github.com/oflaned/taraxa-exporter/internal/config"
)



func main() {
	cfg := config.GetConfig()

	fmt.Println("File paths:")
	for _, path := range cfg.FilePaths {
		fmt.Println(path)
	}
}