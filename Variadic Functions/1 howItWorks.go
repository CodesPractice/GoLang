// -------------------- Variadic Functions -------------

// Only the last parameter of a function can be variadic.
//  ... parameter can accept any number of values for that parameter

/*
	func hello(a int, b ...int) {
	}

	hello(1, 2) //passing "1" to a argument "2" to b...
	hello(5, 6, 7, 8, 9) //passing "5" to a argument and  "6, 7, 8, 9" to b...
*/

package main

import "fmt"

// this function checks whether first number can be found in the list of numbers
// inside the function type of nums is []int - an integer slice.
func find(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if num == v {
			found = true
			fmt.Printf("%d is found at index %d in %d", num, i, nums)
		}
	}
	if !found {
		fmt.Printf("%d is not found in %d", num, nums)
	}
	fmt.Println()

}

func main() {
	// pass first number to num argument and other numbers to nums ...
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
}

/*
	89 is found at index 0 in [89 90 95]
	45 is found at index 2 in [56 67 45 90 109]
	78 is not found in [38 56 98]
	87 is not found in []
*/
