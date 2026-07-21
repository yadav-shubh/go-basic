package main

import "fmt"

func main() {
	var name = make([]int32, 2, 5)
	name = append(name, 1)
	fmt.Print(name)
}
