package employee

import "fmt"

// change to the lowercase for struct type and it's field for prevent the exporting
type employee struct {
	firstName   string
	lastName    string
	totLeaves   int
	leavesTaken int
}

// the New functiond takes fields to create an employee and return once it done.

func New(fn string, ln string, tot int, taken int) employee {
	emp := employee{fn, ln, tot, taken}
	return emp
}

func (e employee) ShowRemainingLeaves() {
	fmt.Printf("%s %s has %d leaves remaining.", e.firstName, e.firstName, e.totLeaves-e.leavesTaken)
}
