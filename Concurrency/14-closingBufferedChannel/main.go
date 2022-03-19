//-----------Closing buffered channels-------------

/*
	It's possible to read data from a already closed buffered channel.
	The channel will return the data that is already written to the channel and
	once all the data has been read, it will return the zero value of the channel.
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)

	fmt.Println()

	/*
		Received: 5, open: true
		Received: 6, open: true
		Received: 0, open: false
	*/

	chr := make(chan int, 4)
	chr <- 3
	chr <- 7
	close(chr)
	for v := range chr {
		fmt.Println("Received:", v)
	}

	/*
		Received: 3
		Received: 7
	*/
}
