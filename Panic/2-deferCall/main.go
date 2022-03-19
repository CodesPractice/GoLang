package main

import "fmt"

func fullName(fn *string, ln *string) {
	defer fmt.Println("Defer call in function")
	if fn == nil {
		panic("runtime error :  first name cannot be nil")
	}
	if ln == nil {
		panic("runtime error : last name cannot be nil")
	}
	fmt.Printf("Full name : %s %s \n", *fn, *ln)
	fmt.Println("return normally from function")
}

func main() {
	defer fmt.Println("Defer call from main ")
	first := "Elon"
	fullName(&first, nil)
	fmt.Println("return normally from main")
}
