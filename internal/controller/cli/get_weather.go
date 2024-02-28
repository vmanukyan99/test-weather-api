package cli

import (
	"fmt"
	"time"

	"github.com/vmanukyan99/test-weather-api/internal/entity"
)

func (c *Cli) getWeather() (*entity.Weather, error) {
	var (
		country   string
		city      string
		dateInput string
	)

	_, err := fmt.Scanln(&country)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&country): %w", err)
	}

	_, err = fmt.Scanln(&city)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&city): %w", err)
	}

	_, err = fmt.Scanln(&dateInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&data): %w", err)
	}

	date, err := time.Parse(time.DateOnly, dateInput)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}

	weather, err := c.uc.GetWeather(country, city, date)
	if err != nil {
		return nil, fmt.Errorf("c.uc.GetWeather: %w", err)
	}

	return weather, nil
}
