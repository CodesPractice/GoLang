package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64 // for atomicity
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
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, " - counter : ", counter)
	}
	wg.Done()
}
