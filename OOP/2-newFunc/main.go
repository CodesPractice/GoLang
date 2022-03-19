//--------------New() function instead of constructors--------------

package main

import "newFunc/employee"

func main() {
	// pass values to the employee package New function.
	// it takes para to make a new employee and return a employee

	emp := employee.New("Sunny", "Wilson", 18, 7)
	emp.ShowRemainingLeaves()
	//Sunny Sunny has 11 leaves remaining.
}
