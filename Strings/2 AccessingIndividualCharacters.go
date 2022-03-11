//-------------Accessing individual characters of a string------------

package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Go Language"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}

/*
Characters: G o   L a n g u a g e
Bytes: 47 6f 20 4c 61 6e 67 75 61 67 65
*/
