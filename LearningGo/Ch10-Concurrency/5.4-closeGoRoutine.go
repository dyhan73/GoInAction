package main

import (
	"fmt"
)

func forLoopBug() {
	a := []int{2, 4, 6, 7, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

func fixForLoopBug() {
	a := []int{2, 4, 6, 7, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		v := v           // solution 1 : shadowing by local variable
		go func(v int) { // solution 2 : using parameter
			ch <- v * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func closeChannel() {
	lc := countTo(4)
	for i := range lc {
		fmt.Println(i)
	}
	v, ok := <-lc
	fmt.Println(v, ok) // 채널이 닫혀도 읽기에서 panic 은 발생하지 않으나 ok 가 false
}

func remainedChannel() {
	for i := range countTo(10) {
		if i > 5 {
			break // channel 은 읽히기 위해 영원히 대기함
		}
		fmt.Println(i)
	}
}

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s): // searcher() 결과 올때까지 블록
			case <-done: // done 채널이 닫혔으면 즉시 제로값 반환, 열린채널에서 읽기는 블록
			}
		}(searcher)
	}
	r := <-result
	close(done) // 여기서 닫으면 고루틴의 select 에서 done 을 읽음 (, )
	return r
}

func main5() {
	//forLoopBug()
	//fixForLoopBug()
	//closeChannel()
	remainedChannel()
}
