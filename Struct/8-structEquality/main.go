// ----------- Structs Equality.............
// Structs are value types and are comparable if each of their fields are comparable.
// Two struct variables are considered equal if their corresponding fields are equal.

package main

import "fmt"

type name struct {
	fname, lname string
}

func main() {

	n1 := name{
		fname: "Steve",
		lname: "Jobs",
	}

	n2 := name{
		fname: "Steve",
		lname: "Jobs",
	}
	n3 := name{
		fname: "Steve",
	}

	if n1 == n2 {
		fmt.Printf("%s and %s are equal", n1, n2)
	} else {
		fmt.Printf("%s and %s are NOT equal", n1, n2)
	}
	fmt.Println()
	if n1 == n3 {
		fmt.Printf("%s and %s are equal", n1, n2)
	} else {
		fmt.Printf("%s and %s are NOT equal", n1, n3)
	}

}
