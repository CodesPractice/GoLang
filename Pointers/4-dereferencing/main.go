//------------- Dereferencing a pointer--------------

/*
	Dereferencing a pointer means
	accessing the value of the variable to which the pointer points.
	*a is the syntax to deference a.
*/

package main

import "fmt"

func main() {
	val := 11
	ptr := &val
	fmt.Println("address of val is", ptr) //address of val is 0xc000012088
	fmt.Println("value of val is", *ptr)  //value of val is 11
	fmt.Println()

	//----------------------

	v := 22
	p := &v
	fmt.Println("address of v is", p) //address of v is 0xc0000120b0
	fmt.Println("value of v is", *p)  //value of v is 22
	*p++
	fmt.Println("new value of v is", v) //new value of v is 23
}
