package employee

import "fmt"

//define strunt and it's field with capital for the export/accessibility

type Employee struct {
	FirstNam    string
	LastName    string
	TotalLeaves int
	LeavesTake  int
}

func (e Employee) ShowRemainingLeaves() {
	fmt.Printf("%s %s has %d leaves remaining. ", e.FirstNam, e.LastName, e.TotalLeaves-e.LeavesTake)
}
