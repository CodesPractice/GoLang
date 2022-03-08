//-------------- Passing a slice to a function-----------

//changes made to the slice inside the function are visible outside the function.
//When a slice is passed to a function, the pointer variable will refer to the same underlying array.

package main

import "fmt"

func substractTwo(number []int) {
	for i := range number {
		number[i] -= 2
	}
}

func main() {

	no := []int{3, 7, 9, 4}
	fmt.Println("Original slice ", no)
	substractTwo(no)
	fmt.Println("Slice after calling the function ", no)
}
