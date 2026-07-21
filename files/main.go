package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with file in go")
	MyFile, err := os.Create(`names.txt`)
	defer MyFile.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("File created successfully")
		MyArray := []string{"Hello", "World", "How", "are", "you"}
		for _, value := range MyArray {
			_, _ = MyFile.WriteString(value + "\n")
		}
		ReadFile(MyFile.Name())
	}
}

func ReadFile(file string) {
	databytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("File contents are(bytes) : \n", databytes)
	fmt.Println("File contents are(string) : \n", string(databytes))
}
