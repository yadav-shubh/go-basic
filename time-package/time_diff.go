package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 200000000; i++ {
		fmt.Print(i, " ")
	}

	end := time.Since(start)
	fmt.Printf("Execution time: %v\n", end)
}
