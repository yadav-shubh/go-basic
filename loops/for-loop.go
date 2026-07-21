package main

import "fmt"

func main() {
	fmt.Println("This is the loop program")
	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//}

	//for {
	//	fmt.Println("Hello")
	//}

	//i := 10
	//for i < 100 {
	//	fmt.Println(i)
	//}

	//for index, value := range []string{"name", "dog", "ram"} {
	//	fmt.Println(index, " : ", value)
	//}

	for index, value := range "name" {
		fmt.Println(index, " : ", value)
	}
}
