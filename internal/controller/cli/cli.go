package cli

import (
	"fmt"
	"log"

	"github.com/vmanukyan99/test-weather-api/internal/usecase"
)

type Cli struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Cli {
	return &Cli{
		uc: uc,
	}
}

func (c *Cli) Start() {
	for {
		info, err := c.getWeather()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
