//-------------Rune ------------

// A rune is a builtin type in Go and it's the alias of int32.
// Rune represents a Unicode code point in Go.

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
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func main() {
	name := "Hello Andrew"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}

/*
Characters: H e l l o   A n d r e w
Bytes: 48 65 6c 6c 6f 20 41 6e 64 72 65 77

String: Señor
Characters: S e ñ o r
Bytes: 53 65 c3 b1 6f 72
*/
