/*
2. Passing the file path as a command line flag
Another way to solve this problem is to pass the file path as a command line argument.
Using the flag package, we can get the file path as input argument from the command line and
then read its contents.
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error :", err)
		return
	}
	fmt.Println("Content of the file : ", string(data))
	//fmt.Println("value of fpath is", *fptr)
}

// PS D:\VIRTUAN\GO\ReadingFiles\1-ReadingEntireFile\filehandling>
// filehandling --fpath=/VIRTUAN/GO/ReadingFiles/1-ReadingEntireFile/filehandling/test.txt
