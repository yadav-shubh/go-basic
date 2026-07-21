package main

import "fmt"

func main() {
	if i := 9; i < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("positive ", i)
	}
}
