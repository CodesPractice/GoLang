package main

import (
	"fmt"
	"math"
)

func circleArea(rad float32) (float32, error) {
	if rad < 0 {
		return 0, fmt.Errorf("Area calculation failed!, radious %0.2f", rad)
	}
	return math.Pi * rad * rad, nil
}

func main() {
	radius := -20.5
	area, err := circleArea(float32(radius))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of the Circle %0.2f", area)
}
