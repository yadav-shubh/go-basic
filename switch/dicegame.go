package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(7)
	switch dice {
	case 1:
		fmt.Println("your dice is 1")
		break
	case 2:
		fmt.Println("your dice is 2")
		break
	default:
		fmt.Println("invalid")

	}
}
