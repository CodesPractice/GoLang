package main

import "fmt"

type SalaryCalculator interface {
	calculatSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

// p implements the interface. Now Permanent struct also calculatSalary type
func (p Permanent) calculatSalary() int {
	return p.basicPay + p.pf
}

// c implements the interface. Now Contract struct also calculatSalary type
func (c Contract) calculatSalary() int {
	return c.basicPay
}

// f implements the interface. Now Freelancer struct also calculatSalary type
func (f Freelancer) calculatSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpenses(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculatSalary()
		//fmt.Printf("%T\n", v)
	}
	fmt.Printf("Total Expense Per Month $ %d", expense)
}

func main() {
	pEmp1 := Permanent{101, 5000, 50}
	pEmp2 := Permanent{105, 3500, 35}

	cEmp1 := Contract{1001, 1250}

	fEmp1 := Freelancer{1002, 60, 100}
	fEmp2 := Freelancer{1003, 40, 80}

	// Permanent and Contract also SalaryCalculator type
	employees := []SalaryCalculator{pEmp1, pEmp2, cEmp1, fEmp1, fEmp2}
	totalExpenses(employees)
}
