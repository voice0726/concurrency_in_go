package main

import (
	"fmt"
	"local/concurrency_in_go/ch4/stream"
	"math/rand"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	intStream := stream.Generate(done, 1, 2, 3, 4)
	pipeline := stream.Multiply(done, stream.Add(done, stream.Multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}

	done2 := make(chan interface{})
	defer close(done2)

	for num := range stream.Take(done, stream.Repeat(done, 1), 10) {
		fmt.Printf("%v", num)
	}

	done3 := make(chan interface{})
	defer close(done3)

	rand := func() interface{} { return rand.Int() }

	for num := range stream.Take(done, stream.RepeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}
