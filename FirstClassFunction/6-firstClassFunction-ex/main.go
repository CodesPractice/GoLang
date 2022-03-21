package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 10}
	r := mulMaps(nums, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}

func mulMaps(n []int, f func(int) int) []int {
	var nums []int
	for _, v := range n {
		nums = append(nums, f(v))
	}
	return nums
}
