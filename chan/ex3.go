package main

import (
	"sync"
	"time"
)

func main() {
	charChan := make(chan string)
	charSlice := []string{"a", "b", "c", "d"}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go doWork("W1", charChan, &wg)
	go doWork("W2", charChan, &wg)

	go func() {
		for i := 0; i < len(charSlice); i++ {
			println("Sending:", charSlice[i])
			select {
			case charChan <- charSlice[i]:
				println("Sent:", charSlice[i])
			case <-time.After(1 * time.Second):
				println("Timeout, skipping:", charSlice[i])
			}
		}

		defer close(charChan)
	}()

	wg.Wait()
}

func doWork(name string, charChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Task Started from", name)
	for char := range charChan {
		time.Sleep(5 * time.Second)
		println(char)
	}
}
