package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	jobs := make(chan int, 10)
	results := make(chan string, 10)

	// Start worker goroutines
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg, results)
	}

	go func() {
		// Send jobs to the workers
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}

func worker(id int, drops <-chan int, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()

	for drop := range drops {
		// Process the drop (for demonstration, we'll just print it)
		fmt.Println("Worker", id, "processing drop", drop)
		result <- fmt.Sprintf("Worker %d processed drop %d", id, drop)
	}
}
