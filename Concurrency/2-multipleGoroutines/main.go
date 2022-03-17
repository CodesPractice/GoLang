//--------------------- Starting multiple Goroutines ----------------

package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

}

func letters() {
	for c := 'a'; c <= 'e'; c++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", c)
	}
}

func main() {
	go numbers()
	go letters()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(" -- main terminated")
}
