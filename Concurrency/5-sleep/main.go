package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello Goroutine is going to sleep! Zzz")
	time.Sleep(4 * time.Second)
	fmt.Println("hello Goroutine awake and going to write to done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main is going to call hello Goroutine !")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}
