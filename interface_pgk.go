package main

import "fmt"

func main() {
	var dog Animal = &Dog{Name: "Dogesh"}
	fmt.Print(dog.Walk())
}

type Animal interface {
	Walk() string
}

type Dog struct {
	Name string
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
