package main

import "fmt"

func main() {
	job := struct {
		id   int
		name string
	}{
		id:   1,
		name: "Shubham",
	}

	fmt.Println(job)
}
