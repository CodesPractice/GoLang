//------------- Type assertion -----------

/*
	Type assertion is used to extract the underlying value of the interface.
	i.(T) is the syntax which is used to get the underlying value of
	interface i whose concrete type is T.
*/

package main

import "fmt"

func assert(i interface{}) {
	s := i.(int)   //get the underlying int value from i
	fmt.Println(s) //-> 56
}
func main() {
	var i1 interface{} = 56
	assert(i1)

	/*
		In following code,  pass i2 of concrete type string to the assert function
		which tries to extract a int value from it.
		This program will panic with the message
	*/

	//var i2 interface{} = "Steven Paul"
	//assert(i2) // panic: interface conversion: interface {} is string, not int

	
}
