package main

import "fmt"

func div60(i int) {
	defer func() {
		fmt.Println("devided by", i)
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
