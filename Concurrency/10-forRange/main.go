/*
	The for range form of the for loop can be used to
	receive values from a channel until it is closed.
*/

package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // send data to channel
	}
	close(ch) //close the channel
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for v := range ch {
		fmt.Println("Received ", v)
	}
}
