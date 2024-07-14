package models

type WeatherTemp struct {
    Low  int    `json:"low"`
    Max  int    `json:"max"`
    Unit string `json:"unit"`
}
