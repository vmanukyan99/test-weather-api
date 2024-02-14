package main

import (
	"github.com/vmanukyan99/test-weather-api/internal/controller/cli"
	"github.com/vmanukyan99/test-weather-api/internal/infrastructure/repository"
	"github.com/vmanukyan99/test-weather-api/internal/usecase"
)

func main() {
	r := repository.New()

	uc := usecase.New(r)

	c := cli.New(uc)

	c.Start()
}
