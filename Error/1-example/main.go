package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt") //For read access.
	if err != nil {                // will be not nil if not found the file
		fmt.Println(err)
		fmt.Println(f)
		return
	}
	fmt.Println(err)
	fmt.Println(f)
	fmt.Println(f.Name(), "opened successfully")
}

//open /file.txt: The system cannot find the file specified.
