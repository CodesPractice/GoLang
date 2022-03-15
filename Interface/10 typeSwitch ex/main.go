//------------- Type switch -----------

package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("\n%s is type of String. ", i.(string)) //Miss MoneyPenny is type of String.
	case int:
		fmt.Printf("\n%d is type of integer. ", i.(int)) //99 is type of integer.
	default:
		fmt.Println("\nUnknown Type!") //Unknown Type!
	}
}

func main() {
	findType("Miss MoneyPenny")
	findType(99)
	findType(56.34)
}
