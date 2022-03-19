package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server 1"

}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"

}

func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	go server1(out1)
	go server2(out2)

	select {
	case o := <-out1:
		fmt.Println(o)
	case o := <-out2:
		fmt.Println(o)
	} //  from server 2

	//--------------------- default case-------------------

	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed when no other case is ready") //default case executed
	}

	//--------------------- nil channels : default case-------------------
	var chNil chan int
	select {
	case <-chNil:
	default:
		fmt.Println("default case executed for nil channels") //default case executed
	}
}
