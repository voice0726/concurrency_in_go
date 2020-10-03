package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		// make a channel inside the function to avoid sending data to it directly
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results // returns the read-only channel
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)

	// chanOwner() <- 1 is error because it returns read-only channel
}
