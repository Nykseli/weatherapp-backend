// Rest api json structs

package main

// Save contains Json data struct from /saveCity/ endpoint
// Data City: contains weather data of single city
type Save struct {
	Data City `json:"data"`
}

// SavedCities contains cities saved on server
// Data []City: Array of saved cities
type SavedCities struct {
	Data []City `json:"data"`
}

// City weather info
// Name string: Name of the city
// Temp float32: Temerature in celcius
type City struct {
	Name string  `json:"cityName"`
	Temp float32 `json:"temperature"`
}
