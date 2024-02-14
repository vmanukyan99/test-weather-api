package usecase

import (
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
)

type Repository interface {
	GetCoordinates(country, city string) (latitude float64, longitude float64, err error)
}

type WeatherAPI interface {
	GetWeather(input *dto.GetWeatherRequest) (*dto.GetWeatherResponse, error)
}

type UseCase struct {
	repository Repository
	weatherAPI WeatherAPI
}

func New(r Repository, w WeatherAPI) *UseCase {
	return &UseCase{
		repository: r,
		weatherAPI: w,
	}
}
