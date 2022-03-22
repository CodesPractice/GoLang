/*
3. Bundling the text file along with the binary

The first step is to install the packr.

Type the following command in the command prompt from the directory to install the package
go get -u github.com/gobuffalo/packr/v2/...
*/

package main

import (
	"fmt"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("fileBox", "../filehandling")
	data, err := box.FindString("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", data)
}
