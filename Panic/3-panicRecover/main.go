package main

import "fmt"

func fullNameRecover() { // 3- execute and recover the panic
	if r := recover(); r != nil { //Here if r is nil then panic did not happened.
		fmt.Println("recovered from ", r)
	}
	//panicking stops and the control returns to the caller : func main
}

func fullName(fn, ln *string) {
	defer fullNameRecover() // 2- found a panic and call the deffer for recover
	if fn == nil {
		panic("runtime error! first name cannot be nil.")
	}
	if ln == nil { // 1- this will executes and panic will occur
		panic("runtime error! last name cannot be nil.")
	}
	fmt.Println("Full Name : ", fn, ln)
	fmt.Println("Returned normally from func fullName")
}

func main() {
	defer fmt.Println("Defer call in func main") //5- finally called the main deffer
	first := "Elon"
	fullName(&first, nil)
	fmt.Println("Returned normally from func main") // 4- no more panic error
}

/*
	recovered from  runtime error! last name cannot be nil.
	Returned normally from func main
	Defer call in func main
*/
