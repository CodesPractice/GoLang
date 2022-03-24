package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go increamenter("foo ") // goroutine #1
	go increamenter("bar ") // goroutine #2
	wg.Wait()
	fmt.Println("Final counter : ", counter)
}

// this is a race condition
// running concurrently and different processes accessing same variable
// overwriting makes mistakes in output
func increamenter(s string) {
	for i := 0; i < 45; i++ {
		x := counter
		x++
		time.Sleep(3 * time.Microsecond)
		counter = x
		fmt.Println(s, i, " - Counter : ", counter)
	}
	wg.Done()
}
