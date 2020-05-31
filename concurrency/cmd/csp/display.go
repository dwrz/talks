package main

import (
	"fmt"
	"time"
)

const (
	title = "TOP TEN ZIP CODES BY HIGHEST CURRENT TEMPERATURE"
)

func display(ztchan chan []*ZipTemp, done chan struct{}) {
	for {
		locations, more := <-ztchan
		if !more {
			done <- struct{}{}
			return
		}

		// Clear the screen.
		print("\033[H\033[2J")

		// Print the title and top ten.
		fmt.Println(title)
		fmt.Printf(
			"Last update: %s.\n\n",
			time.Now().Format(time.RFC1123),
		)

		for i, zt := range locations {
			if zt == nil {
				fmt.Printf("%d. No data\n", i+1)

				continue
			}

			fmt.Printf(
				"%d. %.02f°F %s (%s, %s)\n",
				i+1,
				zt.Temperature,
				zt.ZipCode,
				zt.City,
				zt.State,
			)
		}
	}
}
