//-------------- Length & Capacity -----------

//  Length refers to the number of elements in the slice
//  len(sliceName)

// Capacity of the slice is the number of elements in the underlying array
// starting from the index from which the slice is created.

package main

import "fmt"

func main() {
	arr1 := [7]int{3, 4, 5}
	sl1 := arr1[1:3]                                 //[1][2]
	fmt.Printf("Len %d, Cap %d", len(sl1), cap(sl1)) //Len 2 , Cap 6 (starting form [1] to [7])
	fmt.Println()

	sl2 := []int{2, 3, 5, 7, 2, 9, 1, 0, 3}
	fmt.Printf("Len %d, Cap %d", len(sl2), cap(sl2)) //Len 9, Cap 9
	
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("Original Slice \nlength is  %d and capacity is %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                                           //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing \nlength is", len(fruitslice), "and capacity is", cap(fruitslice))

	/*
		Original Slice
		length is  2 and capacity is 6
		After re-slicing
		length is 6 and capacity is 6
	*/


}
