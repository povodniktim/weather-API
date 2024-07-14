package models

type City struct {
	City      string    `json:"city"`
	Country   string    `json:"country"`
	UpdatedAt int64     `json:"updatedAt"`
	Weather   []Weather `json:"weather"`
}
