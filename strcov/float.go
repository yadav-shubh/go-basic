package main

import (
	"fmt"
	"strconv"
)

func main() {
	float, err := strconv.ParseFloat("1.234", 64)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("parsed")
	}
	fmt.Println(float)
}
