//-----------------Passing a slice to a variadic function-----------------

package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {

	// program will fail with compilation error
	// prog.go:23:10: cannot use nums (type []int) as type int in argument to find
	/*
		nums := []int{89, 90, 95}
		find(89, nums)
	*/

	//  suffix the slice with ellipsis ... and pass
	nums := []int{89, 90, 95}
	find(89, nums...)
}
