//------------- Type assertion -----------

/*
	passing string concrete type to the assert function
	which tries to extract int value from it will panic with the message
	" panic: interface conversion: interface {} is string, not int "

	To solve this we can use v, ok := i.(T)
*/

package main

import "fmt"

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
func main() {
	var s interface{} = 56
	assert(s) //-> 56 true
	var i interface{} = "Miss MoneyPenny"
	assert(i) //-> 0 false
}
