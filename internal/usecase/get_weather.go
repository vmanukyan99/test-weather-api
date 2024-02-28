package usecase

import (
	"fmt"
	"time"

	"github.com/vmanukyan99/test-weather-api/internal/entity"
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
)

func (uc *UseCase) GetWeather(country, city string, date time.Time) (*entity.Weather, error) {
	latitude, longitude, err := uc.repository.GetCoordinates(country, city)
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

	weather := entity.NewWeather(
		fmt.Sprintf("%s %s", country, city),
		date,
		output.Hourly.Temperature2m,
	)

	return weather, nil
}
