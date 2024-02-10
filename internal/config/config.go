package config

import (
	"os"
	"sync"

	"github.com/oflaned/taraxa-exporter/internal/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
    FilePaths []string `yaml:"nodeApi"`
}

var (
    once   sync.Once
    config Config
)

func GetConfig(configPath string) Config {
    once.Do(func() {
		data, err := os.ReadFile(configPath)
		if err != nil {
			utils.Logger.Fatal("failed to read config file")
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			utils.Logger.Fatal("failed to unmarshal config")
		}
	})
    return config
}