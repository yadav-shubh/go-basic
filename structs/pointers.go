package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "Go Language",
		age:  20,
	}

	printStruct(&p1)
	fmt.Printf("Person struct %#v\n", p1)

	p2 := retStruct()
	p2St := *p2
	fmt.Printf("Person struct %#v", p2St)
}
func retStruct() *person {
	return &person{
		name: "Golang Language",
		age:  25,
	}
}

func printStruct(p1 *person) {
	p1.age = 30
	fmt.Printf("Person Age: %v\n", p1)
}
