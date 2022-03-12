//----------------- Declaring pointers----------------

//*T is the type of the pointer variable which points to a value of type T.
// The & operator is used to get the address of a variable.

package main

import "fmt"

func main() {

	val := 55
	var ptr *int = &val

	fmt.Printf("Type of ptr is %T", ptr)     //Type of ptr is *int
	fmt.Println("\nAddress of val is ", ptr) //Address of val is  0xc000012088

}
