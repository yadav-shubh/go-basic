package main

import (
	"fmt"
	"time"
)

func main() {
	go helloWorld("Hello")
	helloWorld("World")
}

func helloWorld(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(str)
	}
}
