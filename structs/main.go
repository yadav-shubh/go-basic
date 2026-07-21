package main

import "fmt"

func main() {
	person := Person{"Shubham", 1}
	fmt.Printf(person.Name)

	type mYUser struct {
		name string
	}
	var user = mYUser{name: "Nmme"}

	fmt.Printf(user.name)
	user.name = "Rakesh"
	fmt.Printf(user.name)

}

func getUsers(name string, age int, email string, status bool) *User {
	return &User{name, age, email, status}
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) getName() {
	fmt.Println("The name of the user is : ", u.Name)
}

type Person struct {
	Name  string
	Class int
}
