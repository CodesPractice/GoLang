package main

import "fmt"

func invalidSliceAccessRecover() {

	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}

}

func invalidSliceAccess() {
	defer invalidSliceAccessRecover()
	sl := []int{2, 6, 1, 9}
	fmt.Println(sl[5])
	fmt.Println("normally returned from func invalidSliceAccess")
}

func main() {
	invalidSliceAccess()
	fmt.Println("normally returned from func main")
}
