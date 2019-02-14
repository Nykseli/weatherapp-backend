package main

import (
	"config"
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi"

	"github.com/gorilla/mux"
)

func init() {

	if len(os.Args) == 2 {
		config.LoadConfig(os.Args[1])
	} else {
		fmt.Print("Usage: ./api <path-to-config>\n")
		os.Exit(1)
	}
}

// our main function
func main() {

	router := mux.NewRouter().StrictSlash(true)
	// HandleFunc functions can be found in rest.go
	router.HandleFunc("/cityWeather/{cityName}", restapi.CityWeather)
	router.HandleFunc("/getCities", restapi.GetCities)
	router.HandleFunc("/saveCity", restapi.SaveCity)
	log.Fatal(http.ListenAndServe(":8000", router))

}
