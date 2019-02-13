package weather

import (
	"bytes"
	"config"
	"encoding/json"
	"net/http"
	"strings"
)

// Info contains relevevant info about the weather
type Info struct {
	Temperature float32 `json:"temperature"`
	CityName    string  `json:"cityName"`
}

// GetCityWeather gets weather info of a city using https://api.openweathermap.org/data/2.5/weather api
// Values are in metric system
func GetCityWeather(cityName string) Info {
	var info Info
	var apiResponse APIResponse
	buf := new(bytes.Buffer)

	requestURL := "https://api.openweathermap.org/data/2.5/weather?units=metric&q=" + cityName + "&APPID=" + config.WeatherAPIKey

	reader := strings.NewReader("")
	request, err := http.NewRequest("GET", requestURL, reader)
	request.Header.Add("Content-type", "application/json")

	_ = err // Ignore error for now

	client := &http.Client{}
	resp, err := client.Do(request)
	_ = err // Ingore error for now

	buf.ReadFrom(resp.Body)
	jsonString := buf.Bytes()
	json.Unmarshal(jsonString, &apiResponse)

	if err != nil {
		//TODO: how to return empty Info?
	}

	info.CityName = apiResponse.Name
	info.Temperature = apiResponse.Main.Temp
	return info
}
