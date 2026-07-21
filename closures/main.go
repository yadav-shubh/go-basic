package main

func main() {
	number := 21
	evenChecker := func() bool {
		flag := true
		if number%2 != 0 {
			flag = false
		}
		return flag
	}
	println(evenChecker())
}
