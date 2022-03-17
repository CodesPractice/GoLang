package main

func main() {
	ch := make(chan int)
	ch <- 5
}

/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        D:/VIRTUAN/GO/concurrency/7-deadlock/main.go:5 +0x31
exit status 2

*/
