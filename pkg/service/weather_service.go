package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type weatherData struct {
	Data mainData `json:"main"`
}

type mainData struct {
	TempKelvin float32 `json:"temp"`
}

type WeatherService interface {
	GetTemperatureForCity(ctx context.Context, cityName string) (CityTemperatureData, error)
}

const API_KEY = "3fc9cb54394cc7ad0010daa12eb9e286"
const WEATHER_URL = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s"

type weatherServiceImpl struct {
	apiKey string
}

func NewWeatherService() WeatherService {
	return &weatherServiceImpl{
		apiKey: API_KEY,
	}
}

func (s *weatherServiceImpl) GetTemperatureForCity(ctx context.Context, cityName string) (CityTemperatureData, error) {
	encodedCityName := url.QueryEscape(cityName)

	url := fmt.Sprintf(WEATHER_URL, encodedCityName, API_KEY)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return CityTemperatureData{}, fmt.Errorf("failed to build request")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return CityTemperatureData{}, fmt.Errorf("failed to query remote API")
	}

	rawWeatherData, err := io.ReadAll(response.Body)
	if err != nil {
		return CityTemperatureData{}, fmt.Errorf("failed to read remote API response")
	}

	var data weatherData
	err = json.Unmarshal(rawWeatherData, &data)
	if err != nil {
		return CityTemperatureData{}, fmt.Errorf("failed to interpret remote API response")
	}

	return generateCityTemperatureData(cityName, data.Data.TempKelvin), nil
}
