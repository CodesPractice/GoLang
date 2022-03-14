// ----- Pointers to a struct -----

package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {

	emp := &Employee{ // emp is the pointer to Employee struct
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}

	fmt.Println("First Name", (*emp).firstName) //First Name Sam
	fmt.Println("Age ", emp.age)                //Age  55
}
