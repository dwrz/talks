package main

import (
	"fmt"
	"sync"

	"github.com/dwrz/talks/concurrency/pkg/weather"
)

const (
	// Hardcoded padding for longest city name, Philadelphia.
	message = "%-12s %.2f°F\n"
)

func printTemperatures(cities []string) {
	for _, city := range cities {
		cityWeather, err := weather.GetCity(city)
		if err != nil {
			fmt.Printf(
				"error getting weather for %s: %v\n", city, err,
			)
			continue
		}

		fmt.Printf(message, city, cityWeather.Temp)
	}
}

func printTemperaturesAsync(cities []string) {
	var weatherReports = make(chan string, len(cities))
	var wg sync.WaitGroup
	wg.Add(len(cities))

	for _, city := range cities {
		go func(city string) {
			defer wg.Done()

			cityWeather, err := weather.GetCity(city)
			if err != nil {
				fmt.Printf(
					"error getting weather for %s: %v\n",
					city, err,
				)
				return
			}

			weatherReports <- fmt.Sprintf(
				message, city, cityWeather.Temp,
			)
		}(city)
	}

	wg.Wait()
	close(weatherReports)

	for report := range weatherReports {
		fmt.Print(report)
	}
}
