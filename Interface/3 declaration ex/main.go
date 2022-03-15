package main

import "fmt"

type Person struct {
	first string
	last  string
}

type secretAgent struct {
	Person
	working bool
}

type human interface {
	speak()
}

// secretAgent has speak() method of human interface
// Therefore secretAgent also type of human
func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last, " - The secret Agent Speaks ")
}

// Person has speak() method of human interface
// Therefore Person also type of human
func (p Person) speak() {
	fmt.Println("I am ", p.first, p.last, " - The Person speaks")
}

func main() {

	// Value can have morethan one type
	sa1 := secretAgent{
		working: true,
		Person: Person{
			first: "Dr.",
			last:  "Yes",
		},
	}

	sa2 := secretAgent{
		working: true,
		Person: Person{
			first: "Dark",
			last:  "Hole",
		},
	}

	p1 := Person{
		first: "Andrew",
		last:  "Ferdinandus",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

}
