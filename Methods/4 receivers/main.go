//------------- Pointer Receivers vs Value Receivers -----------

/*
  The difference between value and pointer receiver is,
  changes made inside a method with a pointer receiver is
  visible to the caller whereas this is not happen in value receiver.

*/

package main

import "fmt"

type Person struct {
	name string
	age  int
}

// assign new value to the name field of value receiver p of type Person
func (p Person) setName(newName string) {
	p.name = newName
}

func (p *Person) setAge(newAge int) {
	p.age = newAge
}

func main() {
	per := Person{"Joy", 30}
	fmt.Println("Bafor change, name is :", per.name) //Bafor change, name is : Joy
	per.setName("Joe")
	fmt.Println("After change, name is :", per.name) //After change, name is : Joy
	fmt.Println()
	fmt.Println("Bafor change, age is :", per.age) //Bafor change, age is : 30
	// (&per).setAge(32)
	per.setAge(32)
	fmt.Println("After change, age is :", per.age) //After change, age is : 32

}
