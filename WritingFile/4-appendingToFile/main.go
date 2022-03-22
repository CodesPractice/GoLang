//------------------Appending a line to a file------------

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644) //open the file in append and write only mode
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine) //add a new line to the file in line
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" bytes writtern successfully!")

}
