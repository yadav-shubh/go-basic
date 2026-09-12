package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := new(sync.WaitGroup)

	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	wg.Add(3)
	// worker creation
	for w := 1; w <= 3; w++ {
		go workerJob(ctx, cancel, jobs, results, wg)
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

func workerJob(ctx context.Context, cancel context.CancelFunc, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	workerID := fmt.Sprintf("Worker-%d", rand.Intn(1000))
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Exiting due to cancellation\n", workerID)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("[%s] Jobs channel closed, exiting\n", workerID)
				return
			}
			
			var result Result
			result.job = job

			// Job 5 triggers fatal error and cancels context
			if job.id == 5 {
				fmt.Printf("[%s] 🔴 FATAL ERROR on job %d - triggering cancellation for all workers\n", workerID, job.id)
				result.err = errors.New("job 5 fatal error - cancelling all workers")
				results <- result
				cancel() // Broadcast cancellation to all workers
				return
			}

			// Simulate failure for every third job (but not job 5)
			if job.id%3 == 0 {
				result.err = errors.New(fmt.Sprintf("job %d processing failed", job.id))
			} else {
				result.output = fmt.Sprintf("%d", job.id)
			}

			results <- result
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
	err    error
}
