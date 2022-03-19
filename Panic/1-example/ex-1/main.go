package main

import "fmt"

func fullName(fname *string, lname *string) {
	if fname == nil {
		panic("runtime error : first name cannot be nil")
	}
	if lname == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("Full name : %s %s \n", *fname, *lname)
	fmt.Println("return normally form function")
}

func main() {
	first := "Elon"
	fullName(&first, nil)
	fmt.Println("return normaly from main function")
}

/*
	panic: runtime error: last name cannot be nil

	goroutine 1 [running]:
	main.fullName(0xc00003c000, 0xc000087f70)
			D:/VIRTUAN/GO/Panic/main.go:10 +0x11a
	main.main()
			D:/VIRTUAN/GO/Panic/main.go:18 +0x35
	exit status 2
*/
