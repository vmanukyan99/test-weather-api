package cli

import (
	"fmt"
	"github.com/vmanukyan99/test-weather-api/internal/entity"
)

func (c *Cli) getWeather() (*entity.Weather, error) {
	var (
		location string
		date     string
	)

	_, err := fmt.Scanln(&location)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&location): %w", err)
	}

	_, err = fmt.Scanln(&date)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&data): %w", err)
	}

	weather, err := c.uc.GetWeather(location, date)
	if err != nil {
		return nil, fmt.Errorf("c.uc.GetWeather: %w", err)
	}

	return weather, nil
}
