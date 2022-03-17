// -----------Channel example program----------

package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello World Goroutine")
	done <- true
}

func numbers(done chan bool) {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d ", i)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go numbers(done)
	go hello(done)
	<-done
	<-done
	fmt.Println("Main Function")
}
