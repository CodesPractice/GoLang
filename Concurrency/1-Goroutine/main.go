//--------------------- Goroutines ----------------

//	Goroutines are functions or methods that run concurrently with other functions or methods.

package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world Goroutine")
}
func main() {
	go hello()                  //starts a new Goroutine.
	time.Sleep(1 * time.Second) //main goroutine is put to sleep for 1 second
	fmt.Println("main function")
}
