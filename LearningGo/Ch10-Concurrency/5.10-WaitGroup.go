package main

import (
	"fmt"
	"sync"
	"time"
)

func simpleWaitGroup() {
	var wg sync.WaitGroup
	fmt.Println("Start", time.Now())
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("doThing1", time.Now())
	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("doThing2", time.Now())
	}()
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("doThing3", time.Now())
	}()
	wg.Wait()
	fmt.Println("Finish", time.Now())
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	fmt.Println(in)
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			out <- processor(<-in)
			//for v := range in {
			//	out <- processor(v)
			//}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

func main() {
	//simpleWaitGroup()

	num := 2
	in := make(chan int, num)
	defer close(in)
	for i := 0; i < num; i++ {
		in <- i
	}
	fmt.Println(in)
	result := processAndGather(in, func(val int) int {
		return val * 2
	}, num)

	fmt.Println(result)

}
