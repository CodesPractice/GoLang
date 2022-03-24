package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/* 	no concurrency
	   	foo()
	   	bar() */

	wg.Add(2) // wait for 2 things

	go foo()  // launches go routine for foo
	go bar()  // launches go routine for bar
	wg.Wait() // wait until waitGroup turn to 0 :  when all are done() this finish the wait
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("foo : ", i)
		time.Sleep(3 * time.Microsecond) // sleep makes concurrency
	}

	wg.Done() // this makes reduce 1 from waitGroup
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar : ", i)
		time.Sleep(20 * time.Microsecond) // sleep makes concurrency
	}

	wg.Done() // this makes reduce 1 from waitGroup
}
