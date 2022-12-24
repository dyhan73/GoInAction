package main

import (
	"errors"
	"fmt"
	"os"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	fmt.Println(opts)
}
func main1() {
	testAddTo()
	multiReturn()
}

func multiReturn() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func divAndRemainderWithNamedReturn(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func testAddTo() {
	namedOrOptionalIsNotOffer()
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...)) // []... => 슬라이스 패킹 해제
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func namedOrOptionalIsNotOffer() {
	// instead using struct
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}
