//-------------- Zero value of Interface -------------

// The zero value of a interface is nil.
// A nil interface has both its underlying value and as well as concrete type as nil

/*
	If we try to call a method on the nil interface,
	the program will panic since the nil interface neither has
	a underlying value nor a concrete type.
*/

package main

import "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
	// d1.Describe()
}
