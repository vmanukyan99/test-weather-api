package entity

import "time"

type Weather struct {
	Location           string
	Date               time.Time
	MaxTemperature     float64
	MinTemperature     float64
	AverageTemperature float64
}

func NewWeather(location string, date time.Time, temperatures []float64) *Weather {
	var (
		maxTemperature, minTemperature = temperatures[0], temperatures[0]
		totalTemperature               float64
	)

	for _, temperature := range temperatures {
		if temperature > maxTemperature {
			maxTemperature = temperature
		}

		if temperature < minTemperature {
			minTemperature = temperature
		}

		totalTemperature += temperature
	}

	averageTemperature := totalTemperature / float64(len(temperatures))

	return &Weather{
		Location:           location,
		Date:               date,
		MaxTemperature:     maxTemperature,
		MinTemperature:     minTemperature,
		AverageTemperature: averageTemperature,
	}
}
