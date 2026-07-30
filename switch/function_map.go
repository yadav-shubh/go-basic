package main

import (
	"fmt"
	"strings"
)

func main() {
	functionMap := make(map[string]func(string))

	functionMap["upper"] = func(s string) {
		fmt.Println(strings.ToUpper(s))
	}

	functionMap["lower"] = func(s string) {
		fmt.Println(strings.ToLower(s))
	}

	need := "upper"
	if f, ok := functionMap[need]; ok {
		f("ssfdfsd")
	}
}
