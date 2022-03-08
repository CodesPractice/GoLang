//------------- Creating slices ---------

// slices do not own any data.
// slices are just a reference to an existing array.

package main

import "fmt"

func main() {

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

	//Another way to create a slice....
	//creates an array and returns a slice reference
	sli1 := []string{"Aa", "Bb", "Cc", "Dd"}
	fmt.Println(sli1) //[Aa Bb Cc Dd]

}
