//------------ Passing pointer to a function--------------

package main

import "fmt"

func change(val *int) {
	//value of a is changed using dereference
	*val = 55
}

func main() {
	v := 58
	fmt.Println("value of v before function call is", v) //value of a before function call is 58
	p := &v
	change(p)
	fmt.Println("value of v after function call is", v) //value of a after function call is 55
}
