package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
)

var key = func() string {
	key := os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	if key == "" {
		panic("missing open weather map api key")
	}

	return key
}()

type Weather struct {
	Temp float64
}

func GetCity(city string) (*Weather, error) {
	params := fmt.Sprintf("appid=%s&q=%s&units=imperial", key, city)
	url := fmt.Sprintf("%s?%s", baseURL, params)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("%s (%d)", body, res.StatusCode)
	}

	var response response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &Weather{Temp: response.Temperatures.Temp}, nil
}

func GetGeo(lat, lng string) (*Weather, error) {
	params := fmt.Sprintf(
		"appid=%s&lat=%s&lon=%s&units=imperial", key, lat, lng,
	)
	url := fmt.Sprintf("%s?%s", baseURL, params)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("%s (%d)", body, res.StatusCode)
	}

	var response response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &Weather{Temp: response.Temperatures.Temp}, nil
}
