package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var nogateMathError = errors.New("Norgate Math :  suqare of negative number")

func main() {
	fmt.Printf("%T \n\n", nogateMathError) //*errors.errorString
	number := -81.2
	sq, err := sqrt(number)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Square root of ", number, " is ", sq)
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, nogateMathError
	}
	fmt.Println()
	return math.Sqrt(num), nil
}

// 2022/03/18 14:27:31 Norgate Math :  suqare of negative number
// exit status 1
