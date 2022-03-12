//------------- Creating pointers using the new function---------------

// The new function takes a type as an argument and
// returns a pointer to a newly allocated zero value of the type passed as argument.

package main

import "fmt"

func main() {

	ptr := new(int) // create a pointer of type int.
	fmt.Printf("ptr value is %d, type is %T, address is %v\n", *ptr, ptr, ptr)
	//ptr value is 0, type is *int, address is 0xc000012088

	*ptr = 85
	fmt.Println("New prt value is", *ptr)
	//New prt value is 85
}
