package weatherapi

import (
	"errors"
	"net/http"
	"time"
)

var ErrGetWeather = errors.New("weather getting is failed")

type WeatherAPI struct {
	host   string
	client http.Client
}

func New(host string) *WeatherAPI {
	return &WeatherAPI{
		client: http.Client{
			Timeout:   30 * time.Second,
			Transport: http.DefaultTransport,
		},
		host: host,
	}
}
