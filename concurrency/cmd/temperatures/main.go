package main

import (
	"fmt"
	"time"
)

var cities = []string{
	"Chicago",
	"Dallas",
	"Houston",
	"Los Angeles",
	"New York",
	"Philadelphia",
	"Phoenix",
	"San Antonio",
	"San Diego",
	"San Jose",
}

func main() {
	now := time.Now()

	printTemperatures(cities)

	fmt.Printf(
		"\nTook %v to get temperatures sequentially.\n\n",
		time.Since(now),
	)

	// Reset the timer.
	now = time.Now()

	printTemperaturesAsync(cities)

	fmt.Printf(
		"\nTook %v to get temperatures concurrently.\n\n",
		time.Since(now),
	)
}
