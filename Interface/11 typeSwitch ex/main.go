//------------- Type switch -----------

package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

type Dog struct {
	name  string
	breed string
	age   int
}

func (p Person) Describe() {
	fmt.Printf("\nHi! I am %s and %d years old.", p.name, p.age)
}

func (d Dog) Describe() {
	fmt.Printf("\n%s is my pet of breed %s, it is %d years old ", d.name, d.breed, d.age)
}

func findType(dis []Describer) {

	for _, v := range dis {

		switch v.(type) {
		case Person:
			fmt.Printf("\n%s is a Person", v.(Person).name)
		case Dog:
			fmt.Printf("\n%s is a Dog", v.(Dog).name)
		default:
			fmt.Printf("\nUnknown type")

		}
	}
}

func main() {
	p := Person{"Naveen", 56}
	d := Dog{"Tina", "Dalmation", 12}
	all := []Describer{p, d}
	p.Describe()
	d.Describe()
	findType(all)
}

/*
	Hi! I am Naveen and 56 years old.
	Tina is my pet of breed Dalmation, it is 12 years old
	Naveen is a Person
	Tina is a Dog
*/
