package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func init() {
	fmt.Println("this is the first init")
}

func init() {
	fmt.Println("this is the second init")
}

func init() {
	fmt.Println("this is the third init")
}
