package main

import "fmt"

func main() {
	//numbers := map[string]int{"one": 1, "two": 2}
	//numbers["three"] = 3
	//for key, value := range numbers {
	//	fmt.Println(key, " = ", value)
	//}

	//names := make(map[int]string)
	//names[1] = "one"

	// check existence of key
	//value, flag := names[2]
	//if flag {
	//	fmt.Println(flag, " = ", value)
	//} else {
	//	fmt.Println("not exist")
	//}

	//delete entry from map
	//delete(names, 2)
	//fmt.Println(names)

	//modify value in map
	//names[1] = "three"
	//fmt.Println(names)

	var data = map[string]string{}

	s, ok := data["name"]
	print(s, ok)

	//for key, value := range data {
	//	fmt.Println(key, value)
	//}

	aMap := map[string]int{}
	aMap = nil
	fmt.Print(aMap)
	aMap["1"] = 1
	fmt.Print(aMap)
}
