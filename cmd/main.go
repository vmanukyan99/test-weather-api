package main

import (
	"github.com/vmanukyan99/test-weather-api/internal/controller/cli"
	weatherapi "github.com/vmanukyan99/test-weather-api/internal/infrastructure/weather_api"
	"github.com/vmanukyan99/test-weather-api/internal/usecase"
)

func main() {
	w := weatherapi.New()

	uc := usecase.New(w)

	c := cli.New(uc)

	c.Start()
}
