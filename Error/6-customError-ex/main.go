package main

import (
	"errors"
	"fmt"
)

func salaryCal(days, rate int) (int, error) {
	if rate == 0 || days == 0 {
		return 0, errors.New("Error in Input Values!")
	}
	return days * rate, nil
}

func main() {
	days := 30
	rate := 220

	sal, err := salaryCal(days, rate)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Salary of the month %d", sal)
}
