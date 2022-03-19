package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Printf("bidirectional channel :\t%T \n", c)  // bidirectional channel : chan int
	fmt.Printf("send only channer  : \t%T \n", cs)   // send only channer  :    chan<- int
	fmt.Printf("receive only channel : \t%T \n", cr) // receive only channel :  <-chan int
}
