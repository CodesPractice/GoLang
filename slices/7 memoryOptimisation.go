//-------------- Memory Optimisation -----------

// Slices hold a reference to arrays and as long as the slice is in memory,
// the array cannot be garbage collected.
// making  a copy of slice will solve this problem and original array can be garbage collected.
package main

import "fmt"

func countries() []string {
	countries := []string{"USA", "UEA", "Canada", "Australia", "India", "Sri Lanka"}
	countriesNeeded := countries[:len(countries)-2]
	// make zero valued slice with required length
	countriesCopy := make([]string, len(countriesNeeded))
	// get a copy of slice and assigned to copied slice
	copy(countriesCopy, countriesNeeded)
	return countries
}

func main() {
	c := countries()
	fmt.Println(c, len(c), cap(c))
	fmt.Print(countries())

}
