package models

type WeatherWind struct {
	Icon          string  `json:"icon"`
	Direction     string  `json:"direction"`
	DirectionLong string  `json:"directionLong"`
	Speed         float64 `json:"speed"`
	Unit          string  `json:"unit"`
}
