package main

import "fmt"

func main() {

	//for i := 10; i > 0; i-- {
	//	defer func(j int) {
	//		fmt.Print(j, " ")
	//	}(i)
	//}

	for i := 3; i > 0; i-- {
		defer fmt.Println(i)
	}

}
