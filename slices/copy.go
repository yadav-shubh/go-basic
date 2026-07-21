package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10, 11}

	copy(a, b)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

}
