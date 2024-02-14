package usecase

import (
	"fmt"
	"github.com/vmanukyan99/test-weather-api/internal/entity"
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
	weatherapi "github.com/vmanukyan99/test-weather-api/internal/infrastructure/weather_api"
	"time"
)

func (uc *UseCase) GetWeather(format string, location string, date time.Time) (*entity.Weather, error) {
	latitude, longitude, err := uc.repository.GetCoordinates(location)
	if err != nil {
		return nil, fmt.Errorf("uc.repository.GetCoordinates: %w", err)
	}

	input := &dto.GetWeatherRequest{
		Latitude:  latitude,
		Longitude: longitude,
		StartDate: date,
		EndDate:   date,
	}

	var host string

	if format == "regular" {
		host = "https://api.open-meteo.com/v1/forecast"
	} else if format == "historical" {
		host = "https://archive-api.open-meteo.com/v1/archive"
	}

	weatherAPI := weatherapi.New(host)

	output, err := weatherAPI.GetWeather(input)
	if err != nil {
		return nil, fmt.Errorf("weatherAPI.GetWeather: %w", err)
	}

	weather := entity.NewWeather(location, date, output.Hourly.Temperature2m)

	return weather, nil
}
