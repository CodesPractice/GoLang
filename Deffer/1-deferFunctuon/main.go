//---------------Defer-----------------

/*
	Defer statement is used to execute a function call
	just before the surrounding function
	where the defer statement is present returns.
*/

package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")
}

func findLargest(nums []int) {
	defer finished()

	fmt.Println("Started finding largest")
	max := nums[0]

	for _, v := range nums {
		if max < v {
			max = v
		}
	}

	fmt.Println("Largest number in", nums, "is", max)

}

func main() {
	nums := []int{3, 1, 7, 4, 0, 2, 9, 3}
	findLargest(nums)
}
