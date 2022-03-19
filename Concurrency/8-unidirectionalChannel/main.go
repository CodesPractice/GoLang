/*
	All is well but what is the point of writing to a send only channel if it cannot be read from!

	This is where channel conversion comes into use.
	It is possible to convert a bidirectional channel to a send only or
	receive only channel but not the vice versa.

*/

package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	//sendch := make(chan<- int)

	sendch := make(chan int) // bidirectional channel sendch is created
	go sendData(sendch)
	fmt.Println(<-sendch)
}
