package main

import (
	"config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CityDataBase is collection of saved cities
var CityDataBase []City

func init() {
	config.LoadConfig()
}

// our main function
func main() {

	router := mux.NewRouter().StrictSlash(true)
	// HandleFunc functions can be found in rest.go
	router.HandleFunc("/cityWeather/{cityName}", CityWeather)
	router.HandleFunc("/getCities", GetCities)
	router.HandleFunc("/saveCity", SaveCity)
	log.Fatal(http.ListenAndServe(":8000", router))

}
