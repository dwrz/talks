package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
	fmt.Printf("took %v\n", time.Since(now))

	now = time.Now()
	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString(fmt.Sprintf("%d\n", i))
	}
	fmt.Println(str.String())
	fmt.Printf("took %v\n", time.Since(now))
}
