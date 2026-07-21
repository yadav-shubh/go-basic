package main

import "fmt"

func main() {
	array := []int{1, 3, 35, 3, 5, 3, 2}
	fmt.Println("array : ", array)
	fmt.Println("slices : ", array[1:4])
}
