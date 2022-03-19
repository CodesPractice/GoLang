package main

import "fmt"

type areaError struct {
	err    string  // error description
	width  float64 // length which caused the error
	length float64 // heright which caused the error
}

func (ae *areaError) Error() string {
	return ae.err
}
func (e *areaError) negativeLength() bool {
	return e.length < 0 // return value is TRUE if less than 0
}

func (e *areaError) negativeWidth() bool {
	return e.width < 0 // return value is TRUE if less than 0
}

func rectArea(length, width float64) (float64, error) {
	err := " "
	if length < 0 && width < 0 {
		err = "Width and Length both are less than zero"
	} else if length < 0 {
		err = "Length is lessthan zero."
	} else if width < 0 {
		err = "Width is lessthan zero."
	}

	if err != " " {
		return 0, &areaError{err, width, length}
	}

	return width * length, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.negativeLength() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.negativeWidth() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}
