package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	charSlice1 := []string{"a", "b", "c", "d", "e"}
	charSlice2 := []string{"a", "b", "c", "d", "e", "f"}

	strChan := make(chan string, 1)

	producerWg := sync.WaitGroup{}
	producerWg.Add(2)

	workerWg := sync.WaitGroup{}
	workerWg.Add(2)

	go producerEx("P1", charSlice1, strChan, &producerWg)
	go producerEx("P2", charSlice2, strChan, &producerWg)

	go func() {
		producerWg.Wait()
		close(strChan)
	}()

	go consumer("C1", strChan, &workerWg)
	go consumer("C2", strChan, &workerWg)

	workerWg.Wait()
}

func consumer(name string, strChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for str := range strChan {
		fmt.Printf("Received: %s by %s\n", str, name)
	}
}

func producerEx(name string, charSlice []string, strChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(charSlice); i++ {
		fmt.Printf("Sending: %s by %s\n", charSlice[i], name)
		strChan <- name + ": " + charSlice[i] + " " + time.Now().Format(time.RFC3339)
		fmt.Printf("Sent: %s by %s\n", charSlice[i], name)
	}
}
