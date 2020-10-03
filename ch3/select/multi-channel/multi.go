package main

import "fmt"

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	// all channels are closed, which means they are readable from main

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		} // both channels are readable, so the select statement randomly determines which case to go
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
