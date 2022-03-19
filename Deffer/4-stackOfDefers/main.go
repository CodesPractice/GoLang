//------------Stack of defers------------

/*
	When a function has multiple defer calls,
	they are pushed on to a stack and executed in Last In First Out (LIFO) order.

*/
package main

import "fmt"

func main() {
	name := "DINUSHIKA"
	fmt.Println("Original string : ", string(name))
	fmt.Print("Reserved string : ")

	for _, v := range []rune(name) {
		defer fmt.Printf("%c ", v)
	}
}
