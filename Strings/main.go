package main

import (
	"fmt"
)

func mutate(s string) string {
	sli := []rune(s)
	sli[0] = 'A'
	return string(sli)
}
func main() {
	h := "hello"
	fmt.Println(mutate(h))
}
