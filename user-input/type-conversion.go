package main

import (
	"bufio"
	"fmt"
	"os"
	strconv "strconv"
	"strings"
)

func main() {
	message := "Enter your name: "
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	number, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if number != 0 {
		fmt.Println("Success : ", number)
	} else {
		fmt.Println("Error : ", err)
	}
}
