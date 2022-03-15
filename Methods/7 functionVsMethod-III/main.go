//------------- Pointer receivers in methods vs Pointer arguments in functions  -----------

// When a function has a pointer argument, it will accept only a pointer argument.
// When a method has a pointer receiver, it will accept both pointer and value receivers.

package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func Area(r *Rectangle) {
	fmt.Printf("Area of Rectangle is %d \n", r.width*r.height)
}

func (r Rectangle) Area() {
	fmt.Printf("Area of Rectangle is %d \n", r.width*r.height)
}

func main() {
	rec := Rectangle{width: 10, height: 15}
	ptr := &rec // ptr is pointer to rec

	ptr.Area() // calling the pointer reciver method
	Area(ptr)  // calling the pointer argument function

	// pointer receiver method accepts both value and pointer receivers
	rec.Area()

	// pinter argument function is not accepting value argument
	// Area(rec)
	// cannot use rec (variable of type Rectangle) as *Rectangle value in argument to Area

}
