//------------- Value receivers in methods vs Value arguments in functions  -----------

// When a function has a value argument, it will accept only a value argument.
// When a method has a value receiver, it will accept both pointer and value receivers.

package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// Method with value receiver of struct type Rectangle
func (r Rectangle) Area() {

	fmt.Printf("Area of the Rectangle is %d", r.width*r.height)
}

// Function with value argument of struct type Rectangle
func Area(r Rectangle) {
	fmt.Printf("Area of the Rectangle is %d", r.width*r.height)
}

func main() {
	rec := Rectangle{
		width:  10,
		height: 20,
	}
	rec.Area() // Caling the Value reciever method
	Area(rec)  //calling the Value argument function

	ptr := &rec // ptr is pointer to rec

	// Area method accepts the pointer receiver too.
	ptr.Area()

	// function is not accepting pointer argument
	// Area(ptr)
	// cannot use ptr (variable of type *Rectangle) as Rectangle value in argument to Area

}
