//---------------Arguments evaluation-------------

/*
	The arguments of a deferred function are evaluated
	when the defer statement is executed,
	 not when the actual function call is done.
*/

package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("value of variable in deferred function", a)
}
func main() {
	a := 25
	defer printA(a)
	a = 10

	fmt.Println("value of variable before deferred function call", a)
}
