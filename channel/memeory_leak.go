package main

import (
	"fmt"
	"sync"
	"time"
)

func leakDemo() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(`sender: before send`)
		for i := 1; i <= 5; i++ {
			fmt.Println(`sender: sending value`, i)
			ch <- i
		}
		fmt.Println(`sender: finished sending without closing the channel`)
	}()

	go func() {
		fmt.Println(`receiver: waiting on range ch`)
		for value := range ch {
			fmt.Println(`receiver: got value`, value)
		}
		fmt.Println(`receiver: exited range`)
	}()

	wg.Wait()
	fmt.Println(`main: sender is done, but receiver is still blocked because channel was never closed`)
	time.Sleep(200 * time.Millisecond)
	fmt.Println(`main: leak demo ends; receiver is still waiting`)
}

func fixedDemo() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(`sender: before send`)
		for i := 1; i <= 5; i++ {
			fmt.Println(`sender: sending value`, i)
			ch <- i
		}
		fmt.Println(`sender: closing channel`)
		close(ch)
	}()

	go func() {
		fmt.Println(`receiver: waiting on range ch`)
		for value := range ch {
			fmt.Println(`receiver: got value`, value)
		}
		fmt.Println(`receiver: exited range normally`)
	}()

	wg.Wait()
	fmt.Println(`main: sender has exited and closed the channel; receiver can now finish`)
	time.Sleep(200 * time.Millisecond)
	fmt.Println(`main: fixed demo completed`)
}

func main() {
	fmt.Println(`=== Leaking version ===`)
	leakDemo()

	fmt.Println()
	fmt.Println(`=== Corrected version ===`)
	fixedDemo()

}
