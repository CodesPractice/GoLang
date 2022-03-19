package main

import (
	"errors"
	"fmt"
	"math"
)

func circleAres(rad float64) (float64, error) {
	if rad < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * rad * rad, nil
}

func main() {
	radius := 20.5

	area, err := circleAres(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
