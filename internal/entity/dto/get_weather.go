package dto

import "time"

type GetWeatherRequest struct {
	Latitude  float64
	Longitude float64
	StartDate time.Time
	EndDate   time.Time
}

type GetWeatherResponse struct {
	Hourly hourlyUnits `json:"hourly"`
}

type hourlyUnits struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}
