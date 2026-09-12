package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	wg.Add(3)
	// worker creation
	for w := 1; w <= 3; w++ {
		go workerJob(jobs, results, wg)
	}

	for w := 1; w <= 10; w++ {
		jobs <- Job{
			id:   w,
			name: fmt.Sprintf("job: %d", w),
		}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		if r.err != nil {
			fmt.Printf("❌ Error: job %d - %v\n", r.job.id, r.err)
		} else {
			fmt.Printf("✅ Job %d result: %s\n", r.job.id, r.output)
		}
	}

}

func workerJob(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		var result Result
		result.job = job

		// Simulate failure for every third job
		if job.id%3 == 0 {
			result.err = errors.New(fmt.Sprintf("job %d processing failed", job.id))
		} else {
			result.output = fmt.Sprintf("%d", job.id)
		}

		results <- result
	}
}

type Job struct {
	id   int
	name string
}

type Result struct {
	job    Job
	output string
	err    error
}
