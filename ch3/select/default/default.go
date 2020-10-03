package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	// no channel is available because both channels are not closed
	default:
		fmt.Printf("in default after %v\n\n", time.Since(start))
	}
}
