// Structure for open weather map api response json for weather endpoint
// https://openweathermap.org/current#current_JSON

package weather

// APIResponse contains the structure of the openweathermap json response
// Coord Coord: location coodrinates.
// Weather []Weather: general info. Weather is always in array
// Base string: Internal parameter
// Main Main: the main info about the weather
// Wind Wind: info about the wind
// Clouds Clouds: info about clouds
// Rain Rain: info about rain
// Snow Snow: info about snow
// Dt int: Time of data calculation, unix, UTC
// Sys Sys: system information
// Id int: City ID
// Name string: City Name
// Cod int: Internal parameter
type APIResponse struct {
	Coord   Coord     `json:"coord"`
	Weather []Weather `json:"weather"`
	Base    string    `json:"base"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Clouds  Clouds    `json:"clouds"`
	Rain    Rain      `json:"rain"`
	Snow    Snow      `json:"snow"`
	Dt      int       `json:"dt"`
	Sys     Sys       `json:"sys"`
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Cod     int       `json:"cod"`
}

// Coord contains location coodrinates.
// Lon float32: location lognitude
// Lat float32: location latitude
type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

// Weather contains general info. Weather is always in array
// ID int: Weather condition id
// Main string: Group of weather parameters (Rain, Snow, Extreme etc.)
// Description: Weather condition within the group
// Icon string: Weather icon id
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main contains the main info about the weather
// Temp float32: Temperature
// Pressure float32: Atmospheric pressure (on the sea level, if there is no sea_level or grnd_level data, hPa
// Humidity float32: Humidity, %
// TempMin float32: Minimum temperature at the moment
// TempMax float32: Maximum temperature at the moment.
// SeaLevel float32: Atmospheric pressure on the sea level, hPa
// GrdnLevel float32: Atmospheric pressure on the ground level, hPa
type Main struct {
	Temp      float32 `json:"temp"`
	Pressure  float32 `json:"pressure"`
	Humidity  float32 `json:"humidity"`
	TempMin   float32 `json:"tmp_min"`
	TempMax   float32 `json:"tmp_max"`
	SeaLevel  float32 `json:"sea_level"`
	GrndLevel float32 `json:"grnd_level"`
}

// Wind contains info about the wind
// Speed float32: Wind speed. Unit Default: meter/sec
// Deg float32: Wind direction, degrees (meteorological)
type Wind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
}

// Clouds contains info about clouds
// All float32: Cloudiness, %
type Clouds struct {
	All float32 `json:"all"`
}

// Rain contains info about rain
// OneH float32: Rain volume for the last 1 hour, mm
// ThreeH float32: Rain volume for the last 3 hour, mm
type Rain struct {
	OneH   float32 `json:"1h"`
	ThreeH float32 `json:"3h"`
}

// Snow contains info about snow
// OneH float32: Snow volume for the last 1 hour, mm
// ThreeH float32: Snow volume for the last 3 hour, mm
type Snow struct {
	OneH   float32 `json:"1h"`
	ThreeH float32 `json:"3h"`
}

// Sys contains system information
// Type int: Internal parameter
// ID int: Internal parameter
// Message float32: Internal parameter
// Country string: Country code (GB, JP etc.)
// Sunrise int: Sunrise time, unix, UTC
// Sunset int: Sunset time, unix, UTC
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}
