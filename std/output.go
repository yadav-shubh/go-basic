package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args

	fmt.Println(args)
	io.WriteString(os.Stdout, strings.Join(args, "-"))
}
