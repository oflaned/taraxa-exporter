package config

import (
	"fmt"
	"os"
	"sync"

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
			panic(fmt.Errorf("failed to read config file: %s", err))
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			panic(fmt.Errorf("failed to unmarshal config: %s", err))
		}
	})
    return config
}