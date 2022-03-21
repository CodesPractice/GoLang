package main

import "fmt"

func simple() func(a, b int) int { // return type is a function declaration
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()         // s is assigned by return type of functuon
	fmt.Println(s(60, 7)) // via s execute the functuon
}
