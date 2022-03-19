package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	w.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mu)
	}
	wg.Wait()
	fmt.Println("Final value of x : ", x)
}

//Final value of x :  1000
