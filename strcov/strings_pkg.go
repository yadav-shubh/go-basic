package main

import (
	"fmt"
	"strings"
)

func main() {
	count := strings.Count("hello", "l")
	fmt.Printf("count of word: %d", count)
	fmt.Println()
	fmt.Printf("contains: %t", strings.Contains("hello", "l"))

	fmt.Println()
	index := strings.Index("hello", "l")
	fmt.Printf("index of word: %d", index)

	fmt.Println()
	fmt.Printf("join: %s", strings.Join([]string{"a", "b", "c"}, ""))

	fmt.Println()
	repeat := strings.Repeat("hello", 3)
	fmt.Printf("repeat len: %s", repeat)

	fmt.Println()
	replace := strings.Replace("hello", "l", "x", 2)
	fmt.Printf("replace len: %s", replace)

	fmt.Println()
	split := strings.Split("hello", "l")
	fmt.Printf("split len: %s", split)
}
