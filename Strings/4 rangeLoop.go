//-------------Accessing individual runes using for range loop ------------

package main

import (
	"fmt"
)

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Señor"
	charsAndBytePosition(name)
}

/*
S starts at byte 0
e starts at byte 1
ñ starts at byte 2
o starts at byte 4
r starts at byte 5
*/
