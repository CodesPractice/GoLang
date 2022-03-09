package main

import "fmt"

func main() {
	empSal := map[string]int{
		"Ann": 12000,
		"Jay": 10000,
		"Joy": 17000,
	}

	fmt.Println("Original map ", empSal)
	edited := empSal
	edited["Joy"] = 13500
	fmt.Println("Original map ", empSal)
}
