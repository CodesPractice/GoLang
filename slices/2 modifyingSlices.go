// ----------  Modifying slices ---------------

// Changes made to the slice will reflect to the original array

package main

import "fmt"

func main() {

	darry := [...]int{5, 10, 15, 20, 25, 30, 35}
	dslice := darry[2:5]                                    // [2][3][4] -> [15 20 25]
	fmt.Println("Array before modifying the slice ", darry) //Array before modifying the slice  [5 10 15 20 25 30 35]
	for i := range dslice {
		dslice[i] += 3
	}
	fmt.Println("Array after modifying the slice ", darry) //Array after modifying the slice   [5 10 18 23 28 30 35]

	// When a number of slices share the same underlying array,
	// the changes that each one makes will be reflected in the array.

	sl1 := darry[:] // creates a slice which contains all elements of the array
	sl2 := darry[:]

	sl1[1] = 00
	fmt.Println("array after modification to slice sl1", darry) //array after modification to slice sl1 [5 0 18 23 28 30 35]
	sl2[3] = 00
	fmt.Println("array after modification to slice sl2", darry) //array after modification to slice sl2 [5 0 18 0 28 30 35]
}
