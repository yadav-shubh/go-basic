package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("Hello ", presentTime.Local())
	futureTime := time.Now()

	const APPNAME = 1
	fmt.Println(APPNAME)
	fmt.Println(futureTime.Sub(presentTime))

	format := presentTime.Format("2006-01-02 15:04:05 PM")
	fmt.Printf("Formatted time: %s\n", format)

	parse, err := time.Parse(time.RFC850, format)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}
