//---------- Declaring and implementing an interface -----------

package main

import (
	"fmt"
	"math"
)

type tank interface {
	Area() float64
	Volume() float64
}

type myValues struct {
	radius float64
	height float64
}

//Area of a tank : 2πr(r + h)
func (m myValues) Area() float64 { //Area methon in the interface
	return 2 * math.Pi * m.radius * (m.height + m.radius)
}

// Volume of a Tank : π r*r h
func (m myValues) Volume() float64 { // Volume methon in the interface
	return math.Pi * m.radius * m.radius * m.height
}

func main() {
	// Accessing elements of
	// the tank interface
	var t tank           // tank interace type variable
	t = myValues{10, 14} // assing myValues to tank type variable
	fmt.Println("Area of tank :", t.Area())
	fmt.Println("Volume of tank:", t.Volume())
}
