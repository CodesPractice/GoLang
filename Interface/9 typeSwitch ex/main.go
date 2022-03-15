//------------- Type switch -----------

/*
	A type switch is used to compare the concrete type of an interface
	against multiple types specified in various case statements.
*/
package main

import "fmt"

func assert(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Printf("Type of %d is interger", i) // -> Type of 56 is interger
	case string:
		fmt.Printf("\nType of %s is string", i) // -> Type of Andy is string
	}

}

func main() {
	var i interface{} = 56
	assert(i)
	var s interface{} = "Miss MoneyPenny"
	assert(s)
}
