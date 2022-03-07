// slices do not own any data.
// slices are just a reference to an existing array.

package main

import "fmt"

func main() {

	// Modifying a slice
	// Changes made to the slice will reflect to the original array
	darry := [...]int{5, 10, 15, 20, 25, 30, 35}
	dslice := darry[2:5] // [2][3][4] -> [15 20 25]
	fmt.Println("Array before modifying the slice ", darry)
	for i := range dslice {
		dslice[i] += 3
	}
	fmt.Println("Array after modifying the slice ", darry)
	//Array before modifying the slice  [5 10 15 20 25 30 35]
	//Array after modifying the slice  [5 10 18 23 28 30 35]

	//---------

	//Another way to create a slice....
	//creates an array and returns a slice reference
	sli1 := []string{"Aa", "Bb", "Cc", "Dd"}
	fmt.Println(sli1) //[Aa Bb Cc Dd]

	//Creating a slice
	arr := [5]int{12, 33, 52, 74, 35}
	var sli []int = arr[1:4] //creates a slice from array index [1] to [3]
	fmt.Println(sli)         //[33 52 74]
	// arr[1:4] -> starting index is : 1 and (length-1) of slice is : 3 ->  [1] [2] [3]

	fmt.Println()
	//creating a slice with all the elements of the arry
	fmt.Println(arr[:]) // [12 33 52 74 35]

	//creating a slice from index 0 to given (length-1) of the arry
	fmt.Println(arr[:3]) // [0][1][2] -> [12 33 52]

	//creating a slice from index 2 to last index of the arry
	fmt.Println(arr[2:]) // [2][3][4] -> [52 74 35]

}
