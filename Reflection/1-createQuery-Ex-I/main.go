package main

import "fmt"

type order struct {
	odrId  int
	custId int
}

func createQuery(o order) string {
	return fmt.Sprintf("insert into order values (%d, %d)", o.odrId, o.custId)
}

func main() {
	o := order{
		odrId:  1234,
		custId: 567,
	}

	fmt.Println(createQuery(o))
}
