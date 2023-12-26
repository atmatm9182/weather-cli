package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Forecast struct {
	Temperature struct {
		Actual    float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Sys struct {
		CountryCode string `json:"country"`
	} `json:"sys"`
	City string `json:"name"`
}

const forecastApiUrl = "https://api.openweathermap.org/data/2.5/weather"

func GetForecast(config *Config) (forecast *Forecast, err error) {
	var u *url.URL
	u, err = url.Parse(forecastApiUrl)
	if err != nil {
		log.Fatal("Could not parse the forecast api url.")
	}

	q := u.Query()

	if config.UseCoords {
		q.Set("lon", strconv.FormatFloat(config.Lon, 'f', -1, 64))
		q.Set("lat", strconv.FormatFloat(config.Lat, 'f', -1, 64))
	} else {
		q.Set("q", config.City)
	}

	q.Set("units", config.Units.String())
	q.Set("appid", config.ApiKey)

	u.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.Get(u.String())
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		err = fmt.Errorf("Weather forecast for city \"%s\" does not exist", config.City)
		return
	case http.StatusUnauthorized:
		err = errors.New("Invalid api key")
		return
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&forecast)

	return
}
