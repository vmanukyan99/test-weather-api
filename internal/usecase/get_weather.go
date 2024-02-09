package usecase

import (
	"fmt"
	"time"

	"github.com/vmanukyan99/test-weather-api/internal/entity"
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
)

func (uc *UseCase) GetWeatherForecast(location string, date time.Time) (*entity.Weather, error) {
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

	output, err := uc.weatherApiForecast.GetWeatherForecast(input)
	if err != nil {
		return nil, fmt.Errorf("uc.weatherAPI.GetWeatherForecast: %w", err)
	}

	weather := entity.NewWeather(location, date, output.Hourly.Temperature2m)

	return weather, nil
}

func (uc *UseCase) GetWeatherArchive(location string, date time.Time) (*entity.Weather, error) {
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

	output, err := uc.weatherApiArchive.GetWeatherArchive(input)
	if err != nil {
		return nil, fmt.Errorf("uc.weatherAPI.GetWeatherArchive: %w", err)
	}

	weather := entity.NewWeather(location, date, output.Hourly.Temperature2m)

	return weather, nil
}
