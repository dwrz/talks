package main

import (
	"sync"
	"time"

	"github.com/dwrz/talks/concurrency/pkg/weather"
)

func getZipTemp(
	wg *sync.WaitGroup, record []string, ztchan chan *ZipTemp,
) {
	defer wg.Done()

	// Ignore malformed records.
	if len(record) < 5 {
		return
	}

	var zt = &ZipTemp{
		City:      record[3],
		Latitude:  record[1],
		Longitude: record[2],
		State:     record[4],
		ZipCode:   record[0],
	}

	// Ignore zip codes we were not able to get the weather for.
	weather, err := weather.GetGeo(zt.Latitude, zt.Longitude)
	if err != nil {
		return
	}

	zt.Temperature = weather.Temp
	zt.Time = time.Now()

	ztchan <- zt
}
