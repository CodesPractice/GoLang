package main

import "fmt"

func digits(num int, ch chan int) {
	for num != 0 {
		digit := num % 10
		ch <- digit
		num /= 10
	}
	close(ch)
}

func calSquare(num int, sqch chan int) {
	sum := 0
	digch := make(chan int)
	go digits(num, digch)

	for v := range digch {
		sum += v * v
	}
	sqch <- sum
}

func calCube(num int, cuch chan int) {
	sum := 0
	digch := make(chan int)
	go digits(num, digch)
	for v := range digch {
		sum += v * v * v
	}
	cuch <- sum
}

func main() {
	number := 3
	sqch := make(chan int)
	cuch := make(chan int)
	go calSquare(number, sqch)
	go calCube(number, cuch)
	s, c := <-sqch, <-cuch
	fmt.Println("Final Output ", s+c)
}
