//-------1. Asserting the underlying struct type and getting more information from the struct fields

package main

import (
	"fmt"
	"os"
)

func main() {
	// open method in os package
	// func open(name string)(file *File, err error )

	/*
		-----------------------------------------------------------------------------------------------
			type PathError struct {
				Op   string
				Path string
				Err  error
			}
			func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
		-----------------------------------------------------------------------------------------------
	*/

	// in here *File is assigned to f and error is assigned to err
	f, err := os.Open("test.txt")

	if err != nil { // if there is an error err will be not nil

		//if pErr, ok := err.(*os.PathError); ok {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			fmt.Println(ok)
			return
		}

		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
