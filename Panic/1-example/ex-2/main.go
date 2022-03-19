package main

import "fmt"

func panicSlice() {
	sl := []int{2, 3, 4}
	fmt.Println(sl[4]) // slice out of index
	fmt.Println("normally returned from function")
}

func main() {
	panicSlice()
	fmt.Println("normally returned from main")
}

/*
panic: runtime error: index out of range [4] with length 3

goroutine 1 [running]:
main.panicSlice()
        D:/VIRTUAN/GO/Panic/ex-2/main.go:7 +0x1d
main.main()
        D:/VIRTUAN/GO/Panic/ex-2/main.go:12 +0x19
exit status 2
*/
