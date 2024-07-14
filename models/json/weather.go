package models

type Weather struct {
	Day   string      `json:"day"`
	Valid string      `json:"valid"`
	Icon  string      `json:"icon"`
	Temp  WeatherTemp `json:"temp"`
	Wind  WeatherWind `json:"wind"`
}
