package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // goroutines are blocked here till the channel receive something or is closed
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Unblocking goroutines...")
	close(begin) // unblock here
	wg.Wait()
}
