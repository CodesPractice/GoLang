//------------- Methods vs Functions -----------
//----------- Methods with the same name -----------

// Methods with the same name can be defined on different types.
// whereas functions with the same names are not allowed.

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  int
	height int
}

type Circle struct {
	redius float64
}

// Area method with value receiver of Rectangle type
func (r Rectangle) Area() int {
	return r.width * r.height
}

// Area mothd with value receiver of Circle type
func (c Circle) Area() float64 {
	return math.Pi * c.redius * c.redius
}

func main() {

	rec := Rectangle{width: 10, height: 20}
	fmt.Printf("Area of the Rectangle %d", rec.Area()) //Area of the Rectangle 200
	cir := Circle{redius: 7}
	fmt.Printf("\nArea of the Rectangle %f", cir.Area()) //Area of the Rectangle 153.938040

}
