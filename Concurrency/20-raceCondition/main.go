package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()

	fmt.Println("Final value of x is ", x)

}

// different value for each run
// Final value of x is  991
// Final value of x is  998
// Final value of x is  984
