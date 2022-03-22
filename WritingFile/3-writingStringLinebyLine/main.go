package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a string slice

	s := []string{"Welcome to the world of Go.", "Go is a compiled language.", "It is easy to learn Go."}

	for _, v := range s {
		fmt.Fprintln(f, v) //The Fprintln function takes a io.writer as parameter and appends a new line
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	/* 	for _, v := range s {
		l, err := f.WriteString(string(v))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

	} */

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" bytes writtern successfully!")

}
