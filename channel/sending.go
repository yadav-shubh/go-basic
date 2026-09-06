package main

import (
	"fmt"
	"sync"
)

func main() {
	multiSender()
}

func multiSender() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	cwg := &sync.WaitGroup{}
	cwg.Add(1)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 10; i <= 20; i++ {
			ch <- i
		}
	}()

	go func() {
		defer cwg.Done()
		for val := range ch {
			fmt.Println("Received: ", val)
		}
	}()

	wg.Wait()
	close(ch)
	cwg.Wait()
}
