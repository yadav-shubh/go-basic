package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stringList := []string{"A", "B", "C", "D", "E"}
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go aka2(ctx, wg, jobs, stringList)
	go majdoor2(wg, jobs)

	wg.Wait()
}

func majdoor2(wg *sync.WaitGroup, jobs <-chan string) {
	defer wg.Done()
	for str := range jobs {
		fmt.Printf("consumer received: %s\n", str)
	}
}

func aka2(ctx context.Context, wg *sync.WaitGroup, jobs chan<- string, stringList []string) {
	defer wg.Done()
	defer close(jobs)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for _, str := range stringList {
		// Step 1: Wait for the rate-limit gate (or context cancellation)
		select {
		case <-ctx.Done():
			fmt.Printf("producer aborted before tick: %s (%v)\n", str, ctx.Err())
			return // Exit immediately!
		case <-ticker.C:
		}

		// Step 2: Rate gate opened; now send with cancellation protection
		select {
		case <-ctx.Done():
			fmt.Printf("producer aborted during send: %s (%v)\n", str, ctx.Err())
			return
		case jobs <- str:
			fmt.Printf("producer sent: %s\n", str)
		}
	}
}
