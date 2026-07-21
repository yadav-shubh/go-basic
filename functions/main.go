package main

import (
	"fmt"
)

func main() {
	//call by value
	a := 10
	b := 20
	//fmt.Printf("befor swapping a : %d and b %d\n", a, b)
	//swap(a, b)
	//fmt.Printf("after swapping a : %d and b %d", a, b)

	//call by reference
	fmt.Printf("befor swapping a : %d and b %d\n", a, b)
	swapByReference(&a, &b)
	fmt.Printf("after swapping a : %d and b %d", a, b)
}

func swap(a int, b int) {
	temp := 0
	temp = a
	a = b
	b = temp
}

func swapByReference(a, b *int) {
	temp := 0
	temp = *a
	*a = *b
	*b = temp
}
