package computer

type Spec struct { // exported struct
	Make  string // exported field
	Price int    // exported field
	model string //unexported field
}
