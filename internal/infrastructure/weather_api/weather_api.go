package weatherapi

import (
	"errors"
	"net/http"
	"time"
)

var ErrGetWeather = errors.New("weather getting is failed")

type WeatherApiForecast struct {
	host   string
	client http.Client
}

func NewForecast(host string) *WeatherApiForecast {
	return &WeatherApiForecast{
		client: http.Client{
			Timeout:   30 * time.Second,
			Transport: http.DefaultTransport,
		},
		host: host,
	}
}

type WeatherApiArchive struct {
	host   string
	client http.Client
}

func NewArchive(host string) *WeatherApiArchive {
	return &WeatherApiArchive{
		client: http.Client{
			Timeout:   30 * time.Second,
			Transport: http.DefaultTransport,
		},
		host: host,
	}
}
