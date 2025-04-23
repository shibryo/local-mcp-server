package infra

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shibryo/local-mcp-server/internal/domain"
)

// Package weather provides a WeatherAPI client
type WeatherAPI struct {
}

// NewWeatherAPI creates a new WeatherAPI instance
func NewWeatherAPI() WeatherAPI {
	return WeatherAPI{}
}

// GetWeather fetches the weather for a given location (e.g. "130010") from the weather API
func (w WeatherAPI) GetWeather(location string) (*domain.Weather, error) {
	resp, err := http.Get("https://weather.tsukumijima.net/api/forecast/city/" + location)
	if err != nil {
		return nil, fmt.Errorf("wether data fetching: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code cheking: %s", resp.Status)
	}

	// Parse the response
	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("weather data parsing: %w", err)
	}
	if len(weatherData.Forecasts) == 0 {
		return nil, fmt.Errorf("no weather data found for location: %s", location)
	}
	forecast := weatherData.Forecasts[0]
	// ドメインモデルに変換
	weather := &domain.Weather{
		Location: domain.Location{
			Area:       weatherData.Location.Area,
			Prefecture: weatherData.Location.Prefecture,
			District:   weatherData.Location.District,
			City:       weatherData.Location.City,
		},
		Condition: forecast.Telop,
		Temperature: domain.Temperature{
			Current: forecast.Temperature.Max.Celsius, // 現在の気温は最高気温を使用
			Max:     forecast.Temperature.Max.Celsius,
			Min:     forecast.Temperature.Min.Celsius,
		},
		Description: weatherData.Description.Text,
	}

	return weather, nil
}

// WeatherResponse represents the response from the weather API
type WeatherResponse struct {
	PublicTime          string      `json:"publicTime"`
	PublicTimeFormatted string      `json:"publicTimeFormatted"`
	PublishingOffice    string      `json:"publishingOffice"`
	Title               string      `json:"title"`
	Link                string      `json:"link"`
	Description         Description `json:"description"`
	Forecasts           []Forecast  `json:"forecasts"`
	Location            Location    `json:"location"`
	Copyright           Copyright   `json:"copyright"`
}

type Description struct {
	PublicTime          string `json:"publicTime"`
	PublicTimeFormatted string `json:"publicTimeFormatted"`
	HeadlineText        string `json:"headlineText"`
	BodyText            string `json:"bodyText"`
	Text                string `json:"text"`
}

type Forecast struct {
	Date         string       `json:"date"`
	DateLabel    string       `json:"dateLabel"`
	Telop        string       `json:"telop"`
	Detail       Detail       `json:"detail"`
	Temperature  Temperature  `json:"temperature"`
	ChanceOfRain ChanceOfRain `json:"chanceOfRain"`
	Image        Image        `json:"image"`
}

type Detail struct {
	Weather string `json:"weather"`
	Wind    string `json:"wind"`
	Wave    string `json:"wave"`
}

type Temperature struct {
	Min TemperatureValue `json:"min"`
	Max TemperatureValue `json:"max"`
}

type TemperatureValue struct {
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
}

type ChanceOfRain struct {
	T00_06 string `json:"T00_06"`
	T06_12 string `json:"T06_12"`
	T12_18 string `json:"T12_18"`
	T18_24 string `json:"T18_24"`
}

type Image struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Link   string `json:"link,omitempty"`
}

type Location struct {
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	District   string `json:"district"`
	City       string `json:"city"`
}

type Copyright struct {
	Title    string     `json:"title"`
	Link     string     `json:"link"`
	Image    Image      `json:"image"`
	Provider []Provider `json:"provider"`
}

type Provider struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Note string `json:"note"`
}
