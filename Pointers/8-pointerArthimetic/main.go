// Go does not support pointer arithmetic

package main

func main() {
	b := [...]int{109, 110, 111}
	p := &b
	p++
}

// this code will throw compilation error
// main.go:6: invalid operation: p++ (non-numeric type *[3]int)
