package main

import (
	"fmt"
	"io"
	"reflect"
)

type MyInt int

func typeAssertion() {
	var i interface{}
	fmt.Println(reflect.TypeOf(i))
	var mine MyInt = 20
	i = mine
	fmt.Println(reflect.TypeOf(i))
	i2 := i.(MyInt)
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2, reflect.TypeOf(i))
	//i3 := i.(string) // panic: interface conversion: interface {} is main.MyInt, not string
	//fmt.Println(i3)
}

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Println("i is nil")
	case int:
		fmt.Println("i is int", j)
	case MyInt:
		fmt.Println("i is MyInt")
	case io.Reader:
		fmt.Println("i is io.Reader")
	case string:
		fmt.Println("i is string")
	case bool, rune:
		fmt.Println("i is bool or rune")
	default:
		fmt.Println("i don't know")
	}
}
func typeSwitch() {
	p1 := "asdf"
	doThings(p1)
	fmt.Println(reflect.TypeOf(p1))
}

func main11() {
	typeAssertion()
	typeSwitch()
}
