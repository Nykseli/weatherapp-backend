package main

import (
	"config"
	"fmt"
	"weather"
)

func init() {
	config.LoadConfig()
}

// our main function
func main() {
	var info weather.Info
	info = weather.GetCityWeather("Turku")
	fmt.Printf("Name of the city: %s\n", info.CityName)
	fmt.Printf("City temperature: %.2f\n", info.Temperature)

}
