package main

import "fmt"

func main() {
	var dog Animal = &Dog{Name: "Dogesh"}
	fmt.Print(dog.Walk())
}

type Animal interface {
	Walk() string
	Eat() string
}

type Dog struct {
	Name string
}

func (a *Dog) Eat() string {
	return "Dog can eat"
}

type Fox struct {
	Name string
}

func (a *Dog) Walk() string {
	return "Dog can walk"
}

func (f *Fox) Walk() string {
	return "Fox can walk"
}
