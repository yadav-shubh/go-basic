package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// write all example related to marshal
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	userList := []User{
		{Name: "Rajesh", Age: 22, Email: "rajesh@gmail.com"},
		{Name: "John", Age: 30, Email: "jogn@gmail.com"},
	}

	jsonBytes, err := json.Marshal(userList)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonBytes))
	userList2 := make([]User, 0, 10)
	err = json.Unmarshal(jsonBytes, &userList2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userList2)
}
