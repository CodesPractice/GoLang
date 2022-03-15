//-------------- Embedding interfaces -------------

/*
	Although go does not offer inheritance,
 	it is possible to create a new interfaces by embedding other interfaces.

	Any type is said to implement EmployeeOperations interface
	if it provides method definitions of embeded interfaces
*/

package main

import "fmt"

type SalaryCalculator interface {
	DisplaSalary()
}

type CalculateLeavesLeft interface {
	CalculateLeavesLeft() int
}

// EmployeeOperation interface created by embedding
// SalaryCalculator and LeaveCalculator interfaces.
type EmployeeOperation interface {
	SalaryCalculator
	CalculateLeavesLeft
}

type Employee struct {
	fname       string
	lname       string
	basicPay    int
	pf          int
	totLeaves   int
	leavesTaken int
}

func (e Employee) DisplaSalary() {
	fmt.Printf("Salary of %s %s is %d", e.fname, e.lname, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totLeaves - e.leavesTaken
}

func main() {
	emp := Employee{
		fname:       "Sam",
		lname:       "Walt",
		basicPay:    5000,
		pf:          400,
		totLeaves:   24,
		leavesTaken: 8,
	}

	var empOpt EmployeeOperation = emp
	empOpt.DisplaSalary()
	fmt.Printf("\nLeaves left %d", empOpt.CalculateLeavesLeft())

	// Salary of Sam Walt is 5400
	// Leaves left 16

}
