package main

import "fmt"

func main() {

	a := func() {
		fmt.Println("This is how func assigned to a variable")
	}
	a()
	fmt.Printf("Tyoe of the variable : %T\n", a)

	func() {
		fmt.Println("This is how anonymous func call by itself")
	}() // just after finishing the func declaration use () call it

	func(fn, ln string) {
		fmt.Println("Full name : ", fn, ln)
	}("Dinushika", "Ranasinghe")
}
