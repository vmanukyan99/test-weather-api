package usecase

import (
	"fmt"
	"github.com/vmanukyan99/test-weather-api/internal/entity"
)

func (uc *UseCase) GetWeather(location, date string) (*entity.Weather, error) {
	weather, err := uc.weatherAPI.GetWeather(location, date)
	if err != nil {
		return nil, fmt.Errorf("uc.weatherAPI.GetWeather: %w", err)
	}

	return weather, nil
}
