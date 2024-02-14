package weatherapi

import (
	"encoding/json"
	"fmt"
	"github.com/vmanukyan99/test-weather-api/internal/entity/dto"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (w *WeatherAPI) GetWeather(input *dto.GetWeatherRequest) (*dto.GetWeatherResponse, error) {
	latitude := strconv.FormatFloat(input.Latitude, 'g', -1, 32)
	longitude := strconv.FormatFloat(input.Longitude, 'g', -1, 32)

	log.Println(latitude)
	log.Println(longitude)

	startDate := input.StartDate.Format(time.DateOnly)
	endDate := input.EndDate.Format(time.DateOnly)

	log.Println(startDate)
	log.Println(endDate)

	req, err := http.NewRequest(http.MethodGet, w.host, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	params := url.Values{}
	params.Add("latitude", latitude)
	params.Add("longitude", longitude)
	params.Add("start_date", startDate)
	params.Add("end_date", endDate)
	params.Add("hourly", "temperature_2m")

	req.URL.RawQuery = params.Encode()

	log.Println(req.URL.String())

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("w.client.Do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrGetWeather
	}

	output := &dto.GetWeatherResponse{}

	err = json.NewDecoder(resp.Body).Decode(output)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder(resp.Body).Decode: %w", err)
	}

	return output, nil
}
