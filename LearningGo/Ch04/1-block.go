package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}
