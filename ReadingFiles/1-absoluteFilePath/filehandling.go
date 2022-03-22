package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// access file from the same location
	// data, err := ioutil.ReadFile("test.txt")

	// 1. Giving absolute file path

	data, err := ioutil.ReadFile("D:/VIRTUAN/GO/ReadingFiles/1-ReadingEntireFile/filehandling/test.txt")

	if err != nil {
		fmt.Println("File reading error !", err)
		return
	}
	fmt.Println("Contents of the file", string(data))
}
