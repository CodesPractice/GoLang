package main

import "struct/employee"

func main() {
	//direct access the Employee struct type of the package employee
	emp := employee.Employee{
		FirstNam:    "Sam",
		LastName:    "Pitt",
		TotalLeaves: 26,
		LeavesTake:  12,
	}
	emp.ShowRemainingLeaves() //Sam Pitt has 14 leaves remaining.
}
