//------------- Declaring a method -----------

// A method is just a function with a special receiver type
// between the func keyword and the method name.

/*
	func (t Type) methodName(parameter list) {
	}
*/

package main

import "fmt"

type Person struct {
	name     string
	salary   int
	currency string
}

// displaySalary() method has Person as the receiver type
func (p Person) displaySalaray() {
	fmt.Printf("Salary of %s is %s %d", p.name, p.currency, p.salary)
}

func main() {
	per := Person{"Andy", 5000, " USD"}
	per.displaySalaray() //per is passing to value receiver and calling the method as well.
}
