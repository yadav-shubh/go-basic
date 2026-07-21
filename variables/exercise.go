package main

import "fmt"

func main() {
	// Fahrenheit into Celsius (C = (F − 32) * 5/9).
	var number float32
	fmt.Println("Enter number:")
	fmt.Scanf("%f", &number)

	fmt.Println(number)

	celsius := (number - 32) * 5 / 9
	fmt.Println(celsius)

	meter := number / 0.3048
	fmt.Println(meter)
}
