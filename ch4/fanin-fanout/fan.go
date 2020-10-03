package main

import (
	"fmt"
	"github.com/voice0726/concurrency_in_go/ch4/stream"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	random := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := stream.ToInt(done, stream.RepeatFn(done, random))
	fmt.Println("Primes:")
	for prime := range stream.Take(done, stream.PrimeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))

	done2 := make(chan interface{})
	defer close(done2)

	start2 := time.Now()

	randIntStream2 := stream.ToInt(done, stream.RepeatFn(done, random))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = stream.PrimeFinder(done, randIntStream2)
	}

	for prime := range stream.Take(done, stream.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start2))
}
