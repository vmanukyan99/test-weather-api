package usecase

import (
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
)

type Repository interface {
	GetCoordinates(location string) (latitude float64, longitude float64, err error)
}

type WeatherApiForecast interface {
	GetWeatherForecast(input *dto.GetWeatherRequest) (*dto.GetWeatherResponse, error)
}

type WeatherApiArchive interface {
	GetWeatherArchive(input *dto.GetWeatherRequest) (*dto.GetWeatherResponse, error)
}

type UseCase struct {
	repository         Repository
	weatherApiForecast WeatherApiForecast
	weatherApiArchive  WeatherApiArchive
}

func New(r Repository, waf WeatherApiForecast, waa WeatherApiArchive) *UseCase {
	return &UseCase{
		repository:         r,
		weatherApiForecast: waf,
		weatherApiArchive:  waa,
	}
}
