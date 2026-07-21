package main

func main() {
	chars := []string{"a", "b", "c"}
	charChan := make(chan string)

	go func() {
		defer close(charChan)
		for _, char := range chars {
			//select {
			//case charChan <- char:
			//}
			charChan <- char
		}
	}()

	for char := range charChan {
		println(char)
	}
}
