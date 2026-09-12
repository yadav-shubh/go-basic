package main

import (
	"errors"
	"fmt"
	"sync"
)

type Job struct {
	id   int
	data string
}

type Result struct {
	job    Job
	output string
	err    error
}

func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		var result Result
		result.job = job

		// Simulate failure for every third job
		if job.id%3 == 0 {
			result.err = errors.New(fmt.Sprintf("job %d processing failed", job.id))
		} else {
			result.output = fmt.Sprintf("processed job %d: %s", job.id, job.data)
		}

		results <- result
	}
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	var wg sync.WaitGroup

	// Start 3 workers
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(jobs, results, &wg)
	}

	// Send 10 jobs
	for i := 1; i <= 10; i++ {
		jobs <- Job{id: i, data: fmt.Sprintf("job-%d", i)}
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		if result.err != nil {
			fmt.Printf("❌ Error: %v\n", result.err)
		} else {
			fmt.Printf("✅ %s\n", result.output)
		}
	}
}
