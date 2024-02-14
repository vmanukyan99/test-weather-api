package usecase

import (
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
)

type Repository interface {
	GetCoordinates(location string) (latitude float64, longitude float64, err error)
}

type WeatherAPI interface {
	GetWeather(input *dto.GetWeatherRequest) (*dto.GetWeatherResponse, error)
}

type UseCase struct {
	repository Repository
}

func New(r Repository) *UseCase {
	return &UseCase{
		repository: r,
	}
}
