//----------Deferred methods----------------

package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.fname, p.lname)
}

func main() {
	p := person{
		fname: "Dinu",
		lname: "Ranasinghe",
	}
	defer p.fullName()
	fmt.Print("Welcome! ")
}

//Welcome! Dinu Ranasinghe
