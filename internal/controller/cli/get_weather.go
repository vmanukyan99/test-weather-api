package cli

import (
	"fmt"
	"github.com/vmanukyan99/test-weather-api/internal/entity"
	"time"
)

func (c *Cli) getWeather() (*entity.Weather, error) {
	var (
		locationInput string
		dateInput     string
	)

	_, err := fmt.Scanln(&locationInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&location): %w", err)
	}

	_, err = fmt.Scanln(&dateInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&data): %w", err)
	}

	date, err := time.Parse(time.DateOnly, dateInput)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}

	weather, err := c.uc.GetWeather(locationInput, date)
	if err != nil {
		return nil, fmt.Errorf("c.uc.GetWeather: %w", err)
	}

	return weather, nil
}
