package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Generator(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func main() {

	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	var number = flag.Int("num", 1000, "size of the matrix")

	flag.Parse()
	runtime.GOMAXPROCS(*numCores)

	ch := make(chan int)
	go Generator(ch)

	startTime := time.Now()
	for i := 0; i < *number; i++ {
		prime := <-ch
		//fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	duration := time.Since(startTime)
	fmt.Printf("completed. Number of prime is %d, Time taken: %v\n", *number, duration)
}
