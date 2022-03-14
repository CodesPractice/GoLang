//------------- Anonymous fields-------------
// The default name of an anonymous field is the name as its type
package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// anonymous fields
type Person struct {
	string // default name of this field is string
	int    // default name of this field is int
}

func main() {
	emp4 := Employee{
		firstName: "Anne",
		lastName:  "Marry",
		age:       25,
		salary:    2300,
	}

	fmt.Println("Full Name : " + emp4.firstName + " " + emp4.lastName) // Full Name : Anne Marry
	fmt.Printf("Age : %d years \nSalary : %d", emp4.age, emp4.salary)  // Age : 25 years
	emp4.salary = 3100                                                 // Salary : 2300
	fmt.Printf("\nNew salary : %d", emp4.salary)                       // New salary : 3100
	fmt.Print("\n \n")

	// ----- Anonymous fields in struct -----
	// The default name of an anonymous field is the name of its type
	p1 := Person{
		string: "Naveen",
		int:    52,
	}
	fmt.Println(p1.string) //Naveen
	fmt.Println(p1.int)    // 52
	fmt.Print("\n \n")

}
