package cli

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/vmanukyan99/test-weather-api/internal/entity"
)

var allowedOptions = []int{1, 2}

func validateOption(option int) (int, error) {
	for i := range allowedOptions {
		if allowedOptions[i] == option {
			return option, nil
		}
	}

	return 0, errors.New("chosed wrong option")
}

func (c *Cli) getWeather() (*entity.Weather, error) {
	fmt.Println("Weather")
	fmt.Println("1 - Forecast")
	fmt.Println("2 - Archive")

	var optionInput string

	_, err := fmt.Scanln(&optionInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&optionInput): %w", err)
	}

	option, err := strconv.Atoi(optionInput)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	validatedOption, err := validateOption(option)
	if err != nil {
		return nil, err
	}

	var (
		locationInput string
		dateInput     string
	)

	fmt.Println("Location:")
	_, err = fmt.Scanln(&locationInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&location): %w", err)
	}

	fmt.Println("Date:")
	_, err = fmt.Scanln(&dateInput)
	if err != nil {
		return nil, fmt.Errorf("fmt.Scanln(&data): %w", err)
	}

	date, err := time.Parse(time.DateOnly, dateInput)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}

	var weather *entity.Weather

	if validatedOption == 1 {
		weather, err = c.uc.GetWeatherForecast(locationInput, date)
		if err != nil {
			return nil, fmt.Errorf("c.uc.GetWeatherForecast: %w", err)
		}
	}

	if validatedOption == 2 {
		weather, err = c.uc.GetWeatherArchive(locationInput, date)
		if err != nil {
			return nil, fmt.Errorf("c.uc.GetWeatherArchive: %w", err)
		}
	}

	return weather, nil
}
