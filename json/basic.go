package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	user := User{
		Name:  "Rajesh",
		Age:   22,
		Email: "email@example.com",
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	jsonObj := fmt.Sprintf(`{"name":"Rajesh","age":22}`)

	marshal, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
}
