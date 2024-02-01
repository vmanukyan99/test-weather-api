package usecase

import (
	"fmt"
	"github.com/vmanukyan99/test-weather-api/internal/entity"
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
	"time"
)

func (uc *UseCase) GetWeather(location string, date time.Time) (*entity.Weather, error) {
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

	output, err := uc.weatherAPI.GetWeather(input)
	if err != nil {
		return nil, fmt.Errorf("uc.weatherAPI.GetWeather: %w", err)
	}

	weather := entity.NewWeather(location, date, output.Hourly.Temperature2m)

	return weather, nil
}
