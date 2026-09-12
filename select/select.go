package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10000000; i++ {
			ch <- i
		}
		close(ch)
	}()

	timeout := time.After(2 * time.Second)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context timeout:", ctx.Err())
			return
		case <-timeout:
			fmt.Println("Timeout: No value received in 2 seconds")
			return
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received value:", val)
		}
	}

}
