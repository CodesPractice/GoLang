package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// before any codes starts executing init is the first code to be executed

func init() {
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo : ", i)
		time.Sleep(3 * time.Microsecond) // pause for 3 miliseconds
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar : ", i)
		time.Sleep(20 * time.Microsecond) // pause for 20 miliseconds
	}
	wg.Done()
}
