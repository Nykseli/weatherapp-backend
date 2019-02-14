package tests

import (
	"bytes"
	"config"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"restapi"
	"testing"
	"weather"

	"github.com/gorilla/mux"
)

var router *mux.Router

func TestMain(m *testing.M) {

	router = mux.NewRouter()
	router.HandleFunc("/cityWeather/{cityName}", restapi.CityWeather).Methods("GET")
	router.HandleFunc("/getCities", restapi.GetCities).Methods("GET")
	router.HandleFunc("/saveCity", restapi.SaveCity).Methods("POST")

	testRun := m.Run()
	os.Exit(testRun)
}

func TestEmptySavedCities(t *testing.T) {

	req, _ := http.NewRequest("GET", "/getCities", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if restapi.CityDataBase != nil {
		t.Error("City data base should be empty when saveCity endpoint is not called")
	}

}

// Try to add 1 item to database with /saveCity endpoint
func TestAddCityToDataBase(t *testing.T) {

	// Make sure database is empty
	restapi.CityDataBase = nil

	// create json body
	save := restapi.Save{}
	city := restapi.City{}
	city.Name = "Turku"
	city.Temp = 1.0
	save.Data = city
	body, _ := json.Marshal(save)

	// request
	req, _ := http.NewRequest("POST", "/saveCity", bytes.NewReader(body))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if len(restapi.CityDataBase) != 1 {
		t.Errorf("restapi.CityDataBase length should be 1. Was %d\n", len(restapi.CityDataBase))
	}

}

// Add 5 cities to database and test /getCities endpoint
func TestGetCities(t *testing.T) {

	// Make sure database is empty
	restapi.CityDataBase = nil

	// Add items to database
	city1 := restapi.City{Name: "Turku", Temp: -2.0}
	city2 := restapi.City{Name: "Oulu", Temp: -10.0}
	city3 := restapi.City{Name: "Helsinki", Temp: 0.0}
	city4 := restapi.City{Name: "Tampere", Temp: 2.0}
	city5 := restapi.City{Name: "Lahti", Temp: 1.0}
	restapi.CityDataBase = []restapi.City{city1, city2, city3, city4, city5}

	req, _ := http.NewRequest("GET", "/getCities", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	// response checking
	var respCity restapi.SavedCities
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &respCity)

	if len(respCity.Data) != 5 {
		t.Errorf("Responce should countain 5 items. Returned %d\n", len(respCity.Data))
	}

}

// Try to get Turku weather info
func TestWeatherApi(t *testing.T) {

	if len(os.Getenv("CONFIG_TEST_PATH")) != 0 {
		config.LoadConfig(os.Getenv("CONFIG_TEST_PATH"))
	} else {
		t.Errorf("Set CONFIG_TEST_PATH enviroment variable to run this test")
		t.FailNow()
	}

	req, _ := http.NewRequest("GET", "/cityWeather/Turku", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var info weather.Info
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &info)

	if info.CityName != "Turku" {
		t.Errorf("City name should be Turku. Was %s\n", info.CityName)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatalf("Expected response code %d. Got %d\n", expected, actual)
	}
}
