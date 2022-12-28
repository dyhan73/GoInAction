package main

import "fmt"

func main() {
	examFailed1()
	x := 5
	fmt.Println(&x)
	failedUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}

func update(px *int) {
	*px = 20
}

func examFailed1() {
	var f *int
	failedUpdate(f)
	fmt.Println(f)
}

func failedUpdate(g *int) {
	fmt.Println(g)
	x := 10
	g = &x
	fmt.Println(g)
}
