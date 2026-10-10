package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	maxConcurrent := 1
	semaphore := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		// Acquire a semaphore slot
		semaphore <- struct{}{}

		go func(i int) {
			defer wg.Done()

			defer func() { <-semaphore }()

			fmt.Println("Processing item", i)

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		}(i)
	}

	wg.Wait()
}
