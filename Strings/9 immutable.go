//----------------Strings are immutable-------------

// Strings are immutable in Go.
// Once a string is created it's not possible to change it.

package main

import (
	"fmt"
)

//any valid unicode character within single quote is a rune
func mutate(s string) string {
	s[0] = 'a' // .\main.go:13:7: cannot assign to s[0] (strings are immutable)
	return s
}
func main() {
	h := "hello"
	fmt.Println(mutate(h))
}
