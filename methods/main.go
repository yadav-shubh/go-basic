package main

import "fmt"

func main() {
	person := Person{Name: "Shubham", age: 1}
	person.details(6)
}

func (p Person) details(i int) int {
	fmt.Println(p)
	fmt.Println("number : ", i)
	return i * i
}

type Person struct {
	Name string
	age  int
}
