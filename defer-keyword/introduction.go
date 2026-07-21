// Go program to illustrate the
// concept of the defer statement
package main

import "fmt"

func mul(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, GeeksforGeeks")
}

func main() {

	// Calling mul() function
	// Here mul function behaves
	// like a normal function
	mul(23, 45)

	// Calling mul()function
	// Using defer keyword
	// Here the mul() function
	// is defer function
	defer mul(23, 56)

	// Calling show() function
	show()
}
