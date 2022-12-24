package main

import "fmt"

func main3() {
	cStyle()
	onlyCondition()
	eternalLoop()
	forRange()
	randomMapIteration()
	stringIterByRune()
	forRangeUsingCopyValues()
	labeledForLoop()
}

func labeledForLoop() {
	samples := []string{"hello", "apple_Ф𨮀!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func forRangeUsingCopyValues() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}

func stringIterByRune() {
	// multi-byte rune skips index number
	samples := []string{"hello", "apple_Ф𨮀!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func cStyle() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func onlyCondition() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func eternalLoop() {
	i := 0
	for {
		fmt.Println("Hello")
		i++
		if i > 10 {
			break
		}
	}
}

func forRange() {
	eventVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range eventVals {
		fmt.Println(i, v)
	}
	for _, v := range eventVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"fred": true, "raul": true, "wilma": true}
	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

func randomMapIteration() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
