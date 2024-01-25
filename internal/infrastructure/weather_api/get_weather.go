package weatherapi

import "github.com/vmanukyan99/test-weather-api/internal/entity"

func (w *WeatherAPI) GetWeather(location, date string) (*entity.Weather, error) {
	return &entity.Weather{Temperature: "высокая"}, nil
}
