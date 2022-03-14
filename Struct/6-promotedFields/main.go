// -------------- promoted fields --------------
// Anonymous struct field in a struct are called promoted fields

package main

import "fmt"

type Address struct { // Address is used as a field type for person's addres
	city  string
	state string
}

// Address is the ANONYMOUS STRUCT FIELD in the Person STRUCT
// city and state in the Address struct will be the promoted fields in this case

type Person struct {
	name    string
	age     int
	Address // type Address
}

func main() {
	p1 := Person{
		name: "Andy",
		age:  55,
		Address: Address{ //  type will be the name for anonymous fields
			city:  "Chicago",
			state: "Illionis",
		},
	}

	fmt.Printf("Name : %s \nAge : %d\n", p1.name, p1.age)
	fmt.Println("City : ", p1.city)   // cidy and state are promoted fields
	fmt.Println("State : ", p1.state) // which can access directily
}
