package main

import (
	"config"
	"encoding/json"
	"log"
	"net/http"
	"weather"

	"github.com/gorilla/mux"
)

func init() {
	config.LoadConfig()
}

// our main function
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cityWeather/{cityName}", CityWeather)
	log.Fatal(http.ListenAndServe(":8000", router))

}

// CityWeather handles /cityWeather/{cityName} endpoint
// Return format {"temperatrue": int, "cityName": string}
// If cityName is empty, no weather data was found. try another city
func CityWeather(w http.ResponseWriter, r *http.Request) {
	var info weather.Info
	params := mux.Vars(r)
	info = weather.GetCityWeather(params["cityName"])

	json.NewEncoder(w).Encode(info)
}
