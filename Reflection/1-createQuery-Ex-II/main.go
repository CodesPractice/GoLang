package main

import (
	"fmt"
	"reflect"
)

type order struct {
	odrId  int
	custId int
}

type employee struct {
	id      int
	name    string
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)

}

func main() {
	o := order{
		odrId:  1234,
		custId: 567,
	}

	e := employee{
		name:    "Andy",
		id:      565,
		address: "Science Park Road, Singapore",
		salary:  90000,
		country: "Singapore",
	}

	createQuery(o)
	createQuery(e)

}
