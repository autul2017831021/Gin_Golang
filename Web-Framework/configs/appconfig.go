package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
}

const path = "/configs.json"

var appConfig *AppConfig

func LoadConfig() *AppConfig {
	var loadOnce sync.Once
	loadOnce.Do(loadConfig)
	return appConfig
}

func loadConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filepath := fmt.Sprintf("%s%s", workingDir, path)
	config, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	if parsingErr := json.NewDecoder(config).Decode(&appConfig); parsingErr != nil {
		panic(parsingErr)
	}
}
