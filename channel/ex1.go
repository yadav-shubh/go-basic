package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//leakDemo1()
	fixLeakDemo1()

	time.Sleep(100 * time.Second)
}

func leakDemo1() {
	ch := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			log.Print("Sending: ", i)
			ch <- i
		}
	}()

	go func() {
		for value := range ch {
			log.Print("Received: ", value)
		}
	}()
}

func fixLeakDemo1() {
	ch := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for val := range ch {
			fmt.Println("Received: ", val)
		}
	}()
}
