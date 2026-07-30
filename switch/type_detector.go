package main

import "fmt"

func main() {
	checkType(4)
}

func checkType(i interface{}) string {
	switch typeOf := i.(type) {
	case int8:
		fmt.Printf("int8: %d", typeOf)
	case int:
		fmt.Printf("int : %d", typeOf)
	default:
		fmt.Printf("Type of parameter: %s", typeOf)
	}
	return ""
}
