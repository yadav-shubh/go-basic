package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Enter your name: "
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(welcome)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
