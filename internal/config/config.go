package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
    FilePaths []string `yaml:"file_paths"`
}

var (
    once   sync.Once
    config Config
)

func initializeConfig() {
	data, err := os.ReadFile("config/config.yaml")
    if err != nil {
        panic(fmt.Errorf("failed to read config file: %s", err))
    }

	err = yaml.Unmarshal(data, &config)
    if err != nil {
        panic(fmt.Errorf("failed to unmarshal config: %s", err))
    }
}

func GetConfig() Config {
    once.Do(initializeConfig)
    return config
}