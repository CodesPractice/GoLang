//------------------Zero value of a pointer

// The zero value of a pointer is nil.

package main

import "fmt"

func main() {
	val := 34
	var ptr *int

	if ptr == nil {
		fmt.Println("ptr is ", ptr) //ptr is  <nil>
		ptr = &val
		fmt.Println("ptr after initialization ", ptr) //ptr after initialization  0xc000012088
	}

}
