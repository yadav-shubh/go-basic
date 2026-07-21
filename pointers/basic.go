package main

import "fmt"

func main() {
	mynumber := 23
	number := &mynumber
	fmt.Println(number)
	fmt.Println(*number)
	fmt.Println(*number + *number)
}
