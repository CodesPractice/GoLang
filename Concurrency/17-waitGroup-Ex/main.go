//----------------WaitGroup Ex-----------

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func numbers() {
	// 4 - this will run befor the termination
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// inform the completion
	wg.Done()
}

func letters() {
	for c := 'a'; c <= 'e'; c++ {
		fmt.Printf("%c ", c)
	}
	fmt.Println()

}

func main() {
	// 1- Prints this line
	fmt.Println("Number of Goroutines : ", runtime.NumGoroutine())

	wg.Add(1) //wait for one thing

	go numbers()
	// 2 - run letters()
	letters()
	// 3 - Print this
	fmt.Println("Number of Goroutines : ", runtime.NumGoroutine())

	// wait untile the Done
	wg.Wait()

	// once got the done terminates the programme

}
