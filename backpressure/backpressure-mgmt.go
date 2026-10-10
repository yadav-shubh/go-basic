package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	jobs := make(chan int, 2)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go aka(ctx, wg, jobs)
	go majdoor(wg, jobs)

	wg.Wait()

}

func aka(ctx context.Context, wg *sync.WaitGroup, jobs chan<- int) {
	defer wg.Done()
	defer close(jobs)
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Print("producer canceled\n")
			return
		case jobs <- i:
			fmt.Printf("[Producer] Sent job : %d\n", i)
		}
	}
}

func majdoor(wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("[Worker] Processed job : %d\n", job)
		time.Sleep(100 * time.Millisecond)
	}
}
