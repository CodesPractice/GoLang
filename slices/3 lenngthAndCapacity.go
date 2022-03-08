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

}
