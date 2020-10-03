package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	func1 := func(done chan interface{}) chan int {
		ch := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				ch <- i
				fmt.Printf("[%v] int %v is sent to channel (%v)\n", time.Since(t), i, ch)
			}
			close(ch)
			fmt.Printf("[%v] Channel (%v) closed\n", time.Since(t), ch)
		}()
		return ch
	}
	func2 := func(done chan interface{}, other chan int) chan int {
		ch := make(chan int)
		go func() {
			for i := range other {
				ch <- i
				fmt.Printf("[%v] int %v is received from channel (%v)\n", time.Since(t), i, other)
			}
			close(ch)
			fmt.Printf("[%v] Channel (%v) closed\n", time.Since(t), other)

		}()
		return ch
	}

	done := make(chan interface{})
	defer close(done)
	for range func2(done, func1(done)) {
	}
	simple()
}

func simple() {
	a := make(chan int, 3)
	a <- 1
	a <- 2
	a <- 3
	close(a)

	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
}
