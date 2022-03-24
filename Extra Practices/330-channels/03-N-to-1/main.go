// write to same channel from many goroutines

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	ch := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 10; i < 100; i += 10 {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Print(" ", v)
	}

}
