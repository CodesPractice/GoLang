//-------------String comparison------------

/*
	The == operator is used to compare two strings for equality.
	If both the strings are equal, then the result is true else it's false.

*/

package main

import (
	"fmt"
)

func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

func main() {
	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)

}

/*
Go and Go are equal
hello and world are not equal
*/
