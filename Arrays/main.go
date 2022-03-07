package main

import "fmt"

//Changes made to the array inside a function, will not be reflected in the original array
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("Within function ", num) //Within function  [55 1 3 1 8]
}

//Print two dimentional arrays
func printarray(a [3][2]string) {
	// use blank idetifier to omit index
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}

func main() {
	//Multidimensional arrays

	nos := [3][2]string{
		{"one", "two"},
		{"three", "four"},
		{"five", "six"}, //this comma is necessary. Otherwise compiler will complain
	}
	printarray(nos) //pass 2d araay to funtion to print

	var animals [3][2]string
	animals[0][0] = "cat"
	animals[0][1] = "dog"
	animals[1][0] = "cow"
	animals[1][1] = "sheep"
	animals[2][0] = "whale"
	animals[2][1] = "shark"
	fmt.Println()
	printarray(animals) //pass 2d araay to funtion to print

	// Iteration an array with RANGE
	arrRange := [...]int{5, 2, 8, 1, 0}
	// range returns both index and value
	// if you want to omit the index use _ (blank identifier) instead of first variable
	// for _ , value := range arrRange
	for index, value := range arrRange {
		fmt.Printf("\nIndex %d value is %d", index, value)
	}
	/* Output of the above
	Index 0 value is 5
	Index 1 value is 2
	Index 2 value is 8
	Index 3 value is 1
	Index 4 value is 0
	*/

	// Iteration an array with loop FOR
	// len(array_name) returns the length of the array
	arrFor := [...]int{2, 4, 1, 5, 6}
	for i := 0; i < len(arrFor); i++ {
		fmt.Printf("\n%dth element of array is %d", i, arrFor[i])
	}

	/* Output of the above
	0th element of array is 2
	1th element of array is 4
	2th element of array is 1
	3th element of array is 5
	4th element of array is 6
	*/

	//Changes  made to the array inside a function, will not be reflected in the original array
	num := [...]int{2, 1, 3, 1, 8}
	fmt.Println("Before passing to function", num) //Before passing to function [2 1 3 1 8]
	changeLocal(num)                               // pass array to the function
	fmt.Println("After passing to function", num)  //After passing to function [2 1 3 1 8]

	// Changes are made to the copied array, not be reflected to the original array
	arrOld := [...]string{"Canada", "USA", "UAE"}
	arrNew := arrOld // a copy old array assigned to new array
	arrNew[0] = "Thai"
	fmt.Print("Old Array is ", arrOld)    //Old Array is [Canada USA UAE]
	fmt.Print("\nNew Arrray is ", arrNew) //New Arrray is [Thai USA UAE]

	// To make a copy,  both arrays should have same length
	// In arrays size is also part of the type
	// following code may generates an error since a nd b are in different sizes

	/*  a := [3]int{5, 78, 8}
	var b [4]int
	b = a
	fmt.Print(a) */

	//... makes the compiler determine the length
	vehicles := [...]string{"Car", "Bus", "Jeep", "Van"}
	fmt.Println(vehicles) //[Car Bus Jeep Van]

	//short hand declaraion of an array
	vegs := [3]string{"Carrot", "Tomato", "Potato"}
	fmt.Println(vegs) //[Carrot Tomato Potato]

	// declaration and the initialization at once
	// It is ok to initialize a few elements without all
	var words = [3]string{"Hello"}
	fmt.Println(words) //[Hello  ]

	// declaration and the initialization at once
	var arrFruits = [3]string{"Apple", "Banana", "Grape"}
	fmt.Println(arrFruits) //[Apple Banana grape]

	// declaring a array with 3 elemnts of integer type
	// should mention the size of the array
	var nums [3]int
	fmt.Println(nums) //[0 0 0]

	// Initialization - assigning values to the array
	nums[0] = 13
	nums[1] = 28
	nums[2] = 41
	fmt.Println(nums) //[13 28 41]

}
