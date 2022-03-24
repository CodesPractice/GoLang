package main

import "fmt"

func main() {

	x := 23
	str := checkInt(x)
	fmt.Println(str)

}

func checkInt(x int) string {
	if x > 10 {
		return "x is greater than 10"
	} else {
		return "x is lessthan 10"
	}

}

// in terminal  go to the path and run following code
// golint ./...

// main.go:16:9: if block ends with a return statement, so drop this else and outdent its block
