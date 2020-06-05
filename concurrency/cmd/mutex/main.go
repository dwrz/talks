package main

import (
	"fmt"
	"sync"
)

func main() {
	var balance int = 100
	// var mu sync.Mutex

	fmt.Println("Initial balance:", balance)

	// Apply concurrent adjustments to the balance.
	// The sum of these adjustments is 0.
	// We should see no changes in balance at the end.
	var adjustments = []int{
		-8, -4, -2, 0, 2, 4, 8,
		-8, -4, -2, 0, 2, 4, 8,
		-8, -4, -2, 0, 2, 4, 8,
		-8, -4, -2, 0, 2, 4, 8,
		-8, -4, -2, 0, 2, 4, 8,
		-8, -4, -2, 0, 2, 4, 8,
	}

	var wg sync.WaitGroup
	wg.Add(len(adjustments))

	for _, adjustment := range adjustments {
		// go = placed onto thread, not the stack.
		go func(adjustment int) {
			defer wg.Done()

			// mu.Lock()

			balance += adjustment

			// mu.Unlock()

		}(adjustment)
	}

	wg.Wait()

	fmt.Println("Final balance:", balance)
}
