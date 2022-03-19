//-------------Buffered Channels -----------------

/*
	It is possible to create a channel with a buffer.
	Sends to a buffered channel are blocked only when the buffer is full.
	Similarly receives from a buffered channel are blocked only when the buffer is empty.
*/

package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan string, 2) // make buffered channel with 2 capacity
	ch <- "Sam"                // write to the channel
	ch <- "Willson"            // write to the channel
	//fmt.Println(<-ch)          // read from channel
	//fmt.Println(<-ch)          // read from channell
	fmt.Println()
	//-------------------------

	chw := make(chan int, 2)
	go write(chw)
	time.Sleep(2 * time.Second)
	for v := range chw {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
