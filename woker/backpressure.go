package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id   int
	name string
}

type Result struct {
	job    Job
	output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("    [Worker %d] START processing job %d\n", id, job.id)
		time.Sleep(2 * time.Second) // Simulate slow processing
		fmt.Printf("    [Worker %d] END processing job %d\n", id, job.id)
		results <- Result{
			job:    job,
			output: fmt.Sprintf("Processed by worker %d", id),
		}
	}
	fmt.Printf("    [Worker %d] Exiting\n", id)
}

func main() {
	fmt.Println("=== Buffer Size 3 ===")
	runWithBufferSize(3)

	fmt.Println("\n=== Buffer Size 1 ===")
	runWithBufferSize(1)
}

func runWithBufferSize(bufferSize int) {
	fmt.Printf("Running with buffer size: %d\n", bufferSize)
	fmt.Printf("2 workers, each takes 2 seconds per job\n")
	fmt.Printf("Sending 10 jobs...\n\n")

	wg := new(sync.WaitGroup)
	jobs := make(chan Job, bufferSize)
	results := make(chan Result, 10)

	// Start 2 workers
	numWorkers := 2
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, wg)
	}

	// Producer: send 10 jobs
	startTime := time.Now()
	go func() {
		for i := 1; i <= 10; i++ {
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("[%.1fs] [Producer] Sending job %d...\n", elapsed, i)

			// This will block if buffer is full and workers are slow
			jobs <- Job{id: i, name: fmt.Sprintf("job-%d", i)}

			elapsed = time.Since(startTime).Seconds()
			fmt.Printf("[%.1fs] [Producer] Job %d sent (took %v to send)\n", elapsed, i, time.Since(startTime))
		}
		fmt.Println("[Producer] All jobs sent, closing channel")
		close(jobs)
	}()

	// Wait for workers to finish
	wg.Wait()
	close(results)

	elapsed := time.Since(startTime).Seconds()
	fmt.Printf("\n✅ Completed in %.1f seconds\n", elapsed)
}
