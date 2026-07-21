package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(n int) (fact int) {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
