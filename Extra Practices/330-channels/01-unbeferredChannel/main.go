package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // unbeffered channel has blocking nature

	go func() { // go routine #1
		for i := 0; i < 10; i++ {
			c <- i // code is blocked until value is taken off
		}
	}()

	go func() { // go routine #2
		for {
			fmt.Println(<-c) // once read the value, block will be released
		}
	}()

	time.Sleep(1 * time.Second) // without sleep time, output will not be visible
}

// go run -race main.go
