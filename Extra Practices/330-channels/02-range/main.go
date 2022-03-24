package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := make(chan rune)

	go func() { // anonymous functuon - self executes
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // once finished the putting values to the channel, close the channel
	}()

	for v := range ch {
		fmt.Println(v)
	}
	//------------------------------------------------
	go letters(ch1) //make new GoRoutine and pass channel to get the data.

	for v := range ch1 {
		fmt.Printf(" %c ", v)
	}

}

func letters(c chan rune) {
	for i := 'a'; i < 'l'; i++ {
		c <- rune(i) // wride data to the channel
	}
	close(c) // close the channel for read data
}
