package usecase

import "github.com/vmanukyan99/test-weather-api/internal/entity"

type WeatherAPI interface {
	GetWeather(location, date string) (*entity.Weather, error)
}

type UseCase struct {
	weatherAPI WeatherAPI
}

func New(w WeatherAPI) *UseCase {
	return &UseCase{
		weatherAPI: w,
	}
}
