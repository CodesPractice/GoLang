//------------- Empty interface -----------

/*
	An interface that has zero methods is called an empty interface
	Since the empty interface has zero methods,
	all types implement the empty interface.
*/

package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("\nType = %T and Value = %v", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)

	strt := struct {
		name string
	}{
		name: "Dinushika R",
	}
	describe(strt)
}

/*
	Type = string and Vlue = Hello World
	Type = int and Vlue = 55
	Type = struct { name string } and Vlue = {Dinushika R}
*/
