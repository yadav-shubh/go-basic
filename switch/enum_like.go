package main

import (
	"fmt"
	"math/rand"
)

type Status int

const (
	Pending Status = iota
	InProgress
	Completed
)

func main() {

	num := Status(rand.Intn(3))
	fmt.Printf("Random number: %d\n", num)
	switch num {
	case Pending:
		fmt.Println("Pending")
	case InProgress:
		fmt.Println("InProgress")
		fallthrough
	case Completed:
		fmt.Println("Completed")
	default:
		panic("unhandled default case")
	}

}
