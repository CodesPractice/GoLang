// CodesPractice by Dinushika Ranasinghe ------------------------

// struct is a user-dfined type, that represnts a group of fields

package main

import "fmt"

//struct declaration of Employee
type Employee struct {
	firstName string
	lastNamr  string
	age       int
	salary    int
}

func main() {

	//-----creating struct specifying fields-----
	emp1 := Employee{
		firstName: "Andrew",
		lastNamr:  "Ferdinandus",
		age:       36,
		salary:    800,
	}
	// -----creating struct without specifying fields -----
	emp2 := Employee{"John", "Paul", 29, 750}

	fmt.Println("Employee 1 : ", emp1) //Employee 1 :  {Andrew Ferdinandus 36 800}
	fmt.Println("Employee 2 : ", emp2) //Employee 2 :  {John Paul 29 750}

	// ----- Creating anonymous structs-----
	// It only creates a new struct variable and does not define any new struct type

	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Joy",
		lastName:  "Nikola",
		age:       32,
		salary:    1400,
	}
	fmt.Println("Employee 3 : ", emp3) //Employee 3 :  {Joy Nikola 32 1400}
}
