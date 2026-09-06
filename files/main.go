package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Working with file in go")
	MyArray := []string{"Hello", "World", "How", "are", "you"}

	MyFile := CreateFile(`names.txt`)
	DeferCloseFile(MyFile)

	info, _ := os.Stat(`names.txt`)
	fmt.Println(info.Name(), info.Size(), info.Mode(), info.ModTime())

	WriteToFile(MyArray, MyFile)
	ReadFile(`names.txt`)
}

func DeferCloseFile(file *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}

func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func WriteToFile(data interface{}, file *os.File) {
	err := os.WriteFile(file.Name(), []byte(fmt.Sprintf("%v", data)), 0644)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}

func ReadFile(file string) {
	databases, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("File contents are(bytes) : \n", databases)
	fmt.Println("File contents are(string) : \n", string(databases))
}
