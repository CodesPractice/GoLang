// Do not pass a pointer to an array as an argument to a function.
// Use slice instead.

package main

import (
	"fmt"
)

// a[x] is shorthand for (*a)[x].
// (*arr)[0] can be replaced by arr[0].

func modifyArray(arr *[3]int) {
	//(*arr)[0] = 90
	arr[0] = 90
}

// idiomatic way of achieving this in Go is use slices.

func modifySlice(sl []int) {
	sl[0] = 100
}

func main() {
	a := [3]int{89, 90, 91}
	modifyArray(&a)
	fmt.Println(a) //[90 90 91]

	//---------

	s := []int{10, 20, 20}
	modifySlice(s[:])
	fmt.Println(s) //[100 20 20]

}
