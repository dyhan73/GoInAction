package main

import (
	"errors"
	"fmt"
	"time"
)

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {
	time.Sleep(3 * time.Second)
	return 5, nil
}

func main9() {
	r, e := timeLimit()
	fmt.Println(r, e)
}
