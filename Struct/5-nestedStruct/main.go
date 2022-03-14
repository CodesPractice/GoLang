// ------------ Nested struct--------------
// Struct contains a field which is another struct -> struct inside struct

package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Person struct {
	fname   string
	lname   string
	age     int
	address Address
}

// Address is the STRUCT FIELD in Person STRUCT

func main() {
	p1 := Person{
		fname: "Andrew",
		lname: "Ferdinandus",
		age:   36,
		address: Address{
			city:  "Ragama",
			state: "Gampaha",
		},
	}

	fmt.Printf("Full Name : %s %s\n", p1.fname, p1.lname)
	fmt.Println("Age :", p1.age)
	fmt.Println("City : ", p1.address.city) // access field of inner struct
	fmt.Println("State : ", p1.address.state)
}
