package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	i := -1

	//switch i {
	//case i < 0:
	//	fmt.Println("i < 0")
	//}

	switch {
	case i < 0:
		println("i < 0")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	num := rand.Intn(5)
	switch {
	case num%2 == 0:
		fmt.Println("Even")
	case num%2 != 0:
		fmt.Println("Odd")
	}
}
