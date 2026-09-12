package main

import (
	"fmt"
	"sync"
	"time"
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
		fmt.Println("result", r)
		time.Sleep(1 * time.Second)
	}

}

func workerJob(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Send job to result: %d\n", job.id)
		results <- Result{
			job:    job,
			output: fmt.Sprintf("%d", job.id),
		}
	}
}

type Job struct {
	id   int
	name string
}

type Result struct {
	job    Job
	output string
}
