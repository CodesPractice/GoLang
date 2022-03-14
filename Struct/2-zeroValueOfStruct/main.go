package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {

	//----- Zero value of a struct -----
	var emp5 Employee
	// strings will be " " and integers will be 0
	fmt.Println("First Name :", emp5.firstName) // First Name :
	fmt.Println("Last Name: ", emp5.lastName)   // Last Name:
	fmt.Println("Age : ", emp5.age)             // Age :  0
	fmt.Println("Salary : ", emp5.salary)       // Salary :  0
	fmt.Print("\n \n")

	// Specify values for some fields and ignore the rest.
	// Inored fields are assigned zero values.

	emp6 := Employee{
		firstName: "Mark",
		lastName:  "Watt",
	}
	fmt.Println("First Name :", emp6.firstName) // First Name : Mark
	fmt.Println("Last Name: ", emp6.lastName)   // Last Name:  Watt
	fmt.Println("Age : ", emp6.age)             // Age :  0
	fmt.Println("Salary : ", emp6.salary)       //Salary :  0
}
