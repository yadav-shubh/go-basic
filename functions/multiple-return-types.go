package main

import "fmt"

func main() {
	area, perimeter := rectangleOperation(20, 20)
	fmt.Printf("Area %d, Perimeter %d\n", area, perimeter)
}

func rectangleOperation(height int, width int) (area int, perimeter int) {
	return width * height, (width + height) * 2
}
