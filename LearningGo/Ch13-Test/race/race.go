package race

import (
	"sync"
)

func getCounter() int {
	var counter int
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				counter++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}
