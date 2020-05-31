package main

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
	"time"
)

// The free tier of the OpenWeatherMap API allows for 60 requests per minute.
// ∴ we limit our requests to 1 per second.
var interval = 1 * time.Second

func processCSV(
	file *os.File, ztchan chan *ZipTemp,
) {
	var (
		limiter = time.Tick(interval)
		wg      sync.WaitGroup
	)

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		<-limiter

		wg.Add(1)
		go getZipTemp(&wg, record, ztchan)
	}

	wg.Wait()
	close(ztchan)
}
