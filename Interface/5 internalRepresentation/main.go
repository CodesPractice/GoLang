//-------------------- Interface internal representation

package main

import "fmt"

type Worker interface {
	work()
}

type Person struct {
	name string
}

// Person struct implements the interface Worker
func (p Person) work() {
	fmt.Printf("%s is workng!", p.name)
}

func describeWorker(w Worker) {
	fmt.Printf("\nInterface is Type %T and value is %v", w, w)
}

func main() {
	p := Person{"Dinushika"}
	var w Worker = p  // Now Concrete type of w is Person
	w.work()          // -> Dinushika is workng! Interface is Type main.Person and value is {Dinushika}
	describeWorker(w) // ->
}
