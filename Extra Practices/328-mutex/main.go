// Mutex -> Mutually Exclusive
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

// mutex variable should be in package scope
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo ")
	go incrementor("Boo ")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 45; i++ {
		time.Sleep(20 * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, " - counter : ", counter)
		mutex.Unlock()
	}
	wg.Done()
}
