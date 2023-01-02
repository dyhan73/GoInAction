package main

import "fmt"

// processChannel 버퍼가 있는 채널 예제
// 정확한 고루틴 개수를 알고 있을 때, 고루틴 개수를 제한하거나 대기중인 작업의 양을 제한할 때
func processChannel() []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func(val int) {
			results <- process(val)
		}(i)
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(val int) int {
	return val * 2
}

func main56() {
	results := processChannel()
	fmt.Println(results)
}
