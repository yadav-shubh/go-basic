package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producer(ctx context.Context, wg *sync.WaitGroup, jobs chan<- int) {
	defer wg.Done()
	defer close(jobs)
	for i := 1; i <= 10; i++ {
		select {
		case jobs <- i:
			fmt.Printf("Producer: sending job %d\n", i)
		case <-ctx.Done():
			fmt.Print("Producer: context canceled, stopping production\n")
			return
		}
	}
}

func consumer(jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Consumer: received job %d\n", job)
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	jobs := make(chan int, 2)

	go producer(ctx, wg, jobs)

	//for job := range jobs {
	//	fmt.Printf("Receiving from Producer: job %d\n", job)
	//}
	go consumer(jobs)

	wg.Wait()
}
