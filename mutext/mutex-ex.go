package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	sync.Mutex
}

func main() {

	cnt := new(Counter)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cnt.Lock()
			cnt.value += i
			defer cnt.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(cnt)
}
