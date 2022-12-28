package main

import "fmt"

func main1() {
	exam1()
	pointerType()
	constHasNoAddress()
}

func constHasNoAddress() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}

	fmt.Println(p)
}
func stringp(s string) *string {
	return &s
}

func pointerType() {
	x := 10
	var pointerToX *int
	pointerToX = &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

	var px = new(int)
	fmt.Println(px == nil)
	fmt.Println(*px)
}

func exam1() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

	z := 5 + *pointerToX
	fmt.Println(z)
}
