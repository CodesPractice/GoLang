//----------------Length vs Capacity------------

/*
	The capacity of a buffered channel is
	the number of values that the channel can hold.
	This is the value we specify when creating the buffered
	channel using the make function.

	The length of the buffered channel is
	the number of elements currently queued in it.
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 4
	fmt.Println("Capacity : ", cap(ch))
	fmt.Println("Length : ", len(ch))
	fmt.Println("Read a value from the Channel : ", <-ch)
	fmt.Println("New Length : ", len(ch))
}
