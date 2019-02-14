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
// You can define the path or set it to empty
// If empty try to get ./.config.json
func LoadConfig(path string) {
	var file *os.File
	if len(path) != 0 {
		file, _ = os.Open(path)

	} else {
		file, _ = os.Open(".config.json")

	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic("Config load failed")
	}

	WeatherAPIKey = conf.APIKey
}
