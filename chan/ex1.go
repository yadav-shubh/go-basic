package main

func main() {
	strChan := make(chan string)

	go func() {
		strChan <- "Hello"
	}()

	str := <-strChan
	println(str)
}
