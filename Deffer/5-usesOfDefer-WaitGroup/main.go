// Defer is used in places where
// a function call should be executed irrespective of the code flow.

package main

import (
	"fmt"
	"sync"
)

type rect struct {
	width int
	heigh int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.heigh < 0 {
		fmt.Println("Height is less than 0")
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Println("Width is less than 0")
		wg.Done()
		return
	}
	fmt.Println("Area of the Rectangle is : ", r.heigh*r.width)
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}

	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		w.Add(1)
		go v.area(&w) // Last in first out
	}
	w.Wait()
	fmt.Println("All go routines finished executing")

}
