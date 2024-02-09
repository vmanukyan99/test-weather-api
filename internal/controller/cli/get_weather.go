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
		format        string
	)

	for {
		fmt.Println("Enter format - regular or historical")
		_, err := fmt.Scanln(&format)
		if err != nil {
			return nil, fmt.Errorf("fmt.Scanln(&format): %w", err)
		}

		if format == "regular" || format == "historical" {
			break // Выход из цикла, если введено правильное значение
		} else {
			fmt.Println("Invalid format. Please enter 'regular' or 'historical'.")
		}
	}

	fmt.Println("Enter location - for ex. Moscow")
	_, err := fmt.Scanln(&locationInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&location): %w", err)
	}

	fmt.Println("Enter date - in format 2024-02-09")
	_, err = fmt.Scanln(&dateInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&data): %w", err)
	}

	date, err := time.Parse(time.DateOnly, dateInput)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}

	weather, err := c.uc.GetWeather(format, locationInput, date)
	if err != nil {
		return nil, fmt.Errorf("c.uc.GetWeather: %w", err)
	}

	return weather, nil
}
