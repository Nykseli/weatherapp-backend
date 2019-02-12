package config

import (
	"encoding/json"
	"os"
)

// WeatherAPIKey is api key from .config.json
// If nil, config has failed
var WeatherAPIKey string

// Config contains json fromat for .config.json
type Config struct {
	APIKey string `json:"apiKey"`
}

// LoadConfig loads config from .config.json and load the values to global variables
func LoadConfig() {
	file, _ := os.Open(".config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic("Config load failed")
	}

	WeatherAPIKey = conf.APIKey
}
