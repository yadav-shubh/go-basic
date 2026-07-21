package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	switch strings.TrimSpace(readString) {
	case "Shubham":
		fmt.Println("Shubham")
	case "Hitesh":
		fmt.Println("Hitesh")
	default:
		fmt.Println("Something went wrong")
	}
}
