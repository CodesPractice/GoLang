//-------------- Multidimensional slices-----------

// Similar to arrays
package main

import "fmt"

func main() {

	pls := [][]string{
		{"Cat", "Dog"},
		{"Cow", "Sheep"},
		{"Whale", "Shark"},
	}

	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}

	/*
		Cat Dog
		Cow Sheep
		Whale Shark
	*/

}
