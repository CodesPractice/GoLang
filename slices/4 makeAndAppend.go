//-------------- Creating a slice using Make-----------

// make([]T, len)  can be used to create a slice by passing the type and length to make function
// this will return 0 value elemts slice

package main

import "fmt"

func main() {
	i := make([]int, 5)
	fmt.Println(i) //[0 0 0 0 0]

	//-------------- Add element to slice using append -----------

	// Slices are dynamic and new elements can be add to the slice using append
	// Changes made to the slices will reflect to the original array

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Printf("%d - Length \t%d - Capacity \t Original slice", len(cars), cap(cars))
	//3 - Length      3 - Capacity     Original slice

	cars = append(cars, "Toyota") // capacity of the new slice is become twice the old slice
	fmt.Printf("\n%d - Length \t%d - Capacity \t After Append 1", len(cars), cap(cars))
	// 4 - Length      6 - Capacity     After Append

	cars = append(cars, "Zuzuki")
	fmt.Printf("\n%d - Length \t%d - Capacity \t After Append 2", len(cars), cap(cars))
	//5 - Length      6 - Capacity     After Append 2

	// append one slice to another using the ...
	fmt.Println()
	vegs := []string{"Aloo", "Tomato", "Carrot", "Onion"}
	fruits := []string{"Banana", "Mango", "Apple"}
	food := append(vegs, fruits...)
	fmt.Print(food) //[Aloo Tomato Carrot Onion Banana Mango Apple]
}
