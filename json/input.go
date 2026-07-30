package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	fmt.Print("Start writing\n")
	if scanner.Scan() {
		text = scanner.Text()
	}
	fmt.Println(text)

	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	var user User
	err := json.Unmarshal([]byte(text), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", user)
}
