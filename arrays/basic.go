package main

import "fmt"

func main() {
	//var arr [10]int
	//fmt.Println(arr[0])

	//arr := []string{"ram", "shyam"}
	//for index, value := range arr {
	//	fmt.Println(index, " : ", value)
	//}

	arr := [][]string{
		{
			"ram", "raam",
		}, {
			"shyam",
		}}
	for index, value := range arr {
		for index1, value1 := range value {
			fmt.Println(index, " : ", index1, " : ", value1)
		}
	}

	fmt.Println(len(arr))

	if i := 1.0; i < 0 {
		fmt.Println("if : ", i)
	} else {
		fmt.Println("else : ", i)
	}

	//var count int32 = 10
	//var value float32 = 100
	//
	//fmt.Println(value + count)

}
