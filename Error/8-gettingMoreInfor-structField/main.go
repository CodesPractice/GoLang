package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (ae *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", ae.radius, ae.err)
}

func calculateArea(rad float64) (float64, error) {
	if rad < 0 {
		return 0, &areaError{err: "\nradius is negative", radius: rad}
	}
	return math.Pi * rad * rad, nil
}

func main() {
	radius := -20.0
	area, err := calculateArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			fmt.Println(err)

			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}
