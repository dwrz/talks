package main

func rank(ztchan chan *ZipTemp, ttchan chan []*ZipTemp) {
	var topTen = make([]*ZipTemp, 10)

	for {
		zt, more := <-ztchan
		if !more {
			close(ttchan)
			return
		}

		var newItem bool
		for i, tt := range topTen {
			if tt != nil && tt.Temperature > zt.Temperature {
				continue
			}

			// We've found a zip code with a higher temperature.
			// Insert it into the top ten.
			topTen = append(topTen, nil)
			copy(topTen[i+1:], topTen[i:])
			topTen[i] = zt

			// Drop the last value.
			topTen = topTen[:10]

			newItem = true
			break
		}

		// If the top ten hasn't changed, process the next zip code.
		if !newItem {
			continue
		}

		// If the top ten has changed, message the display goroutine.
		// Send a copy, to prevent race conditions.
		var newTopTen = make([]*ZipTemp, 10)

		for i, tt := range topTen {
			if tt == nil {
				continue
			}
			newTopTen[i] = tt
		}

		ttchan <- newTopTen
	}
}
