// Go program to illustrate how
// to create an anonymous Goroutine
package main

import (
	"fmt"
	"time"
)

// Main function
func main() {

	fmt.Println("Welcome!! to Main function")

	// Creating Anonymous Goroutine
	go func() {
		fmt.Println("Welcome!! to GeeksforGeeks")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}
