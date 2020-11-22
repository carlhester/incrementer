package main

import (
	"fmt"
	"time"
)

func main() {
	n := make(chan int)

	go func() {
		val := 0
		for {
			val++
			n <- val
		}
	}()

	for {
		// every time we read from n, we get an incremented value
		fmt.Println(<-n)
		time.Sleep(1 * time.Second)
	}
}
