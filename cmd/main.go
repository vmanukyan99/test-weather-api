package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/vmanukyan99/test-weather-api/config"
	"github.com/vmanukyan99/test-weather-api/internal/controller/cli"
	"github.com/vmanukyan99/test-weather-api/internal/infrastructure/repository"
	weatherapi "github.com/vmanukyan99/test-weather-api/internal/infrastructure/weather_api"
	"github.com/vmanukyan99/test-weather-api/internal/usecase"
	"log"
)

func main() {
	conf := &config.Config{}

	err := cleanenv.ReadEnv(conf)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := repository.New(conf.DataBase)
	if err != nil {
		log.Fatalln(err)
	}

	w := weatherapi.New("https://api.open-meteo.com/v1")

	uc := usecase.New(r, w)

	c := cli.New(uc)

	c.Start()
}
