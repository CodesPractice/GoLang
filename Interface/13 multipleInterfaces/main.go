// ------------- Implementing multiple interfaces -------------

// A type can implement more than one interface.

package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	fname       string
	lname       string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("Salary of  %s %s is %d ", e.fname, e.lname, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	emp := Employee{
		fname:       "Andrew",
		lname:       "Ferdinandus",
		basicPay:    6000,
		pf:          200,
		totalLeaves: 24,
		leavesTaken: 9,
	}
	//  Hence Employee type implements both interfaces, it is type of both interface
	// emp.DisplaySalary()
	// fmt.Printf("\nLeaves Left : %d", emp.CalculateLeavesLeft())

	var sc SalaryCalculator = emp
	var lc LeaveCalculator = emp

	sc.DisplaySalary()
	fmt.Printf("\nLeaves Left : %d", lc.CalculateLeavesLeft())

	// Salary of  Andrew Ferdinandus is 6200
	// Leaves Left : 15
}
