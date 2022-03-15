//------------- Methods of anonymous struct fields -----------

/*
	Methods belonging to anonymous fields of a struct
	can be called as if they belong to the structure
	where the anonymous field is defined.
*/

package main

import "fmt"

type Address struct {
	city  string
	state string
}

func (a Address) fullAddress() {
	fmt.Printf("Full address is %s, %s", a.city, a.state)
}

type Person struct {
	fname   string
	lname   string
	Address // anonymous struct field (type without name)
}

func main() {

	per := Person{
		fname: "Alice",
		lname: "Brown",
		Address: Address{
			city:  "Brisbane",
			state: "Queensland",
		},
	}

	// Address struct is a field of per
	// accessing fullAddress method of Address struct
	per.fullAddress()
}
