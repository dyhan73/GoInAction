package main

import (
	"fmt"
	"runtime"
)

func panicCase() {
	ch1 := make(chan int, 1) // 버퍼 없으면 패닉 떨어짐
	ch2 := make(chan int, 1)
	go func() {
		v := 1
		ch1 <- v // 채널에 쓸땐 상대쪽에서 읽을때까지 일시중지, 쓰기와 읽기는 sync 되어야 하는 듯
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v
	v2 := <-ch1
	fmt.Println(v, v2)
}

func avoidDeadlockWithSelect() {
	// get function name
	counter, file, line, success := runtime.Caller(0)
	fmt.Println(counter, file, line, success)
	fmt.Println(runtime.FuncForPC(counter).Name())

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}

func main4() {
	panicCase()
	avoidDeadlockWithSelect()
}
