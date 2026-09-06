package main

import (
	"fmt"
	"sync"
	"time"
)

const number int = 1_0_000_00

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(number)
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("Hello from goroutine %d\n", i)
		}()
	}
	wg.Wait()
	end := time.Now()

	fmt.Printf("Elapsed time: %s\n", end.Sub(start))
}
