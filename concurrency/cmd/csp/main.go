package main

import (
	"log"
	"os"
	"time"
)

type ZipTemp struct {
	City        string
	Latitude    string
	Longitude   string
	State       string
	Temperature float64
	Time        time.Time
	ZipCode     string
}

func main() {
	file, err := os.Open("./zipcodes-50.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set up channels.
	var (
		done   = make(chan struct{})
		ztchan = make(chan *ZipTemp)
		ttchan = make(chan []*ZipTemp)
	)

	// Start the rank and display goroutines.
	go rank(ztchan, ttchan)
	go display(ttchan, done)

	// Process the CSV.
	processCSV(file, ztchan)

	// Wait for the last display output.
	<-done
}
