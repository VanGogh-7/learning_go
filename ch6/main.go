package main

import (
	"fmt"
)

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {

	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(pointerX, pointerY, pointerZ)

	p := person{
		FirstName:  "John",
		MiddleName: pointerZ,
		LastName:   "Smith",
	}
	fmt.Println(p)
}
