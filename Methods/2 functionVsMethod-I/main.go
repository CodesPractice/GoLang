//------------- Methods vs Functions -----------

/*
	func (p Person) displaySalaray() {
		fmt.Printf("Salary of %s is %s %d", p.name, p.currency, p.salary)
	}
*/

// Above displaySalary() method converted to function with Employee as parameter

package main

import "fmt"

type Person struct {
	name     string
	salary   int
	currency string
}

func displaySalaray(p Person) {
	fmt.Printf("Salary of %s is %s %d", p.name, p.currency, p.salary)
}

func main() {

	per := Person{"Joe", 4000, "USD"}
	displaySalaray(per) // passing Person as a argument to the func

}
