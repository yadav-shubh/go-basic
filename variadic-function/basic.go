package main

import (
	"fmt"
	"strings"
)

func main() {
	sumAll("string", "string")
	sumAll()
}

func sumAll(x ...string) {
	fmt.Println(strings.Join(x, "-"))
}
