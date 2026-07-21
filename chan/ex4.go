package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	charSlice1 := []string{"a", "b", "c", "d", "e"}
	charSlice2 := []string{"a", "b", "c", "d", "e", "f"}

	strChan := make(chan string, 3)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go producer("P1", charSlice1, strChan, &wg)
	go producer("P2", charSlice2, strChan, &wg)

	go func() {
		wg.Wait()
		close(strChan)
	}()

	for str := range strChan {
		fmt.Printf("Received: %s\n", str)
		time.Sleep(1 * time.Second)
		fmt.Printf("Processing: %s\n", str)
	}

}

func producer(name string, charSlice []string, strChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(charSlice); i++ {
		fmt.Printf("Sending: %s by %s\n", charSlice[i], name)
		strChan <- name + ": " + charSlice[i] + " " + time.Now().Format(time.RFC3339)
		fmt.Printf("Sent: %s by %s\n", charSlice[i], name)
	}
}
