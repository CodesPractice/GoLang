package main

import "fmt"

type add func(a, b int) int
type sub func(a, b int) int
type mul func(a, b int) int
type div func(a, b int) int

func main() {
	x, y := 10, 5
	var a add = func(a, b int) int {
		return a + b
	}

	var s sub = func(a, b int) int {
		return a - b
	}

	var m mul = func(a, b int) int {
		return a * b
	}

	var d div = func(a, b int) int {
		return a / b
	}

	fmt.Println("Addition \t", a(x, y))
	fmt.Println("Substraction \t", s(x, y))
	fmt.Println("Multiplication \t", m(x, y))
	fmt.Println("Division \t", d(x, y))
}
