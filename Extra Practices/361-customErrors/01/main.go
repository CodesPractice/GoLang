package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func Sqrt(num float64) (sq float64, err error) {
	if num < 0 {
		return 0, errors.New("Norgate Math :  suqare of negative number")
	}
	return math.Sqrt(num), nil
}

func main() {
	n1 := -79.0
	sq, err := Sqrt(n1)

	if err != nil {
		//fmt.Println(err)
		log.Fatalln(err)
		return
	}
	fmt.Println("Square root of the ", n1, " is : ", sq)

}

//Square root of the  79  is :  8.888194417315589
//Norgate Math :  suqare of negative number
