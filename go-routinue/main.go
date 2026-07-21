package main

import (
	"fmt"
	"time"
)

func main() {
	printMessage("Renuka")
	go printMessage("Shubham")
}

func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(message + " " + fmt.Sprint(i))
	}
}
