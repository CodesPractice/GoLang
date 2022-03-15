//--------------Implementing interfaces using pointer receivers vs value receivers-----

// value receiver accepts both value and pointer receives,
// but poiter receiver not accepting value receives

package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

type Addresss struct {
	city  string
	state string
}

func (p Person) Describe() {
	fmt.Printf("\n%s is %d years old.", p.name, p.age)
}

func (a *Addresss) Describe() {
	fmt.Printf("\nFull address is %s, %s", a.city, a.state)
}

func displayType(i interface{}) {
	switch v := i.(type) {
	case Person:
		fmt.Printf("\n%s is Person type", v.name)
	case Addresss:
		fmt.Printf("\n%s, %s is Address Type ", v.city, v.state)
	default:
		fmt.Println("Anonymous Type")

	}
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 58}
	d1 = &p1
	d1.Describe()

	p2 := Person{"Andy", 33}
	d1 = p2
	d1.Describe()

	var d2 Describer
	a1 := Addresss{"Los Angeles", "California"}

	/*
		compilation error if the following line is uncommented

		" cannot use a (type Address) as type Describer
		in assignment: Address does not implement
		Describer (Describe method has pointer receiver)
	*/

	//d2 = a
	d2 = &a1
	d2.Describe()
	a2 := Addresss{"Phoenix", "Arizona"}
	//d2 = &a2
	d2.Describe()

	fmt.Print("\n \n")
	displayType(p1)
	displayType(p2)
	displayType(a1)
	displayType(a2)
}

/*
	Sam is 58 years old.
	Andy is 33 years old.
	Full address is Los Angeles, California
	Full address is Los Angeles, California

	Sam is Person type
	Andy is Person type
	Los Angeles, California is Address Type
	Phoenix, Arizona is Address Type

*/
