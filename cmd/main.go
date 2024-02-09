package main

import (
	"github.com/vmanukyan99/test-weather-api/internal/controller/cli"
	"github.com/vmanukyan99/test-weather-api/internal/infrastructure/repository"
	weatherapi "github.com/vmanukyan99/test-weather-api/internal/infrastructure/weather_api"
	"github.com/vmanukyan99/test-weather-api/internal/usecase"
)

func main() {
	r := repository.New()

	waf := weatherapi.NewForecast("https://api.open-meteo.com/v1")
	waa := weatherapi.NewArchive("https://archive-api.open-meteo.com/v1")

	uc := usecase.New(r, waf, waa)

	c := cli.New(uc)

	c.Start()
}
