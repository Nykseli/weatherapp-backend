// Rest api functions

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"weather"

	"github.com/gorilla/mux"
)

// CityWeather handles /cityWeather/{cityName} endpoint
// Return format {"temperatrue": int, "cityName": string}
// If cityName is empty, no weather data was found. try another city
func CityWeather(w http.ResponseWriter, r *http.Request) {
	// Allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var info weather.Info

	params := mux.Vars(r)
	info = weather.GetCityWeather(params["cityName"])

	json.NewEncoder(w).Encode(info)
}

// GetCities returns all the city info saved on the server
func GetCities(w http.ResponseWriter, r *http.Request) {
	// Allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var cities SavedCities
	cities.Data = CityDataBase

	json.NewEncoder(w).Encode(cities)
}

// SaveCity recieves post request body in Save struct json format
func SaveCity(w http.ResponseWriter, r *http.Request) {
	// Allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var save Save
	buf := new(bytes.Buffer)

	buf.ReadFrom(r.Body)
	jsonString := buf.Bytes()

	json.Unmarshal(jsonString, &save)

	CityDataBase = append(CityDataBase, save.Data)

}
