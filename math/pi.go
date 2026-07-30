package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	pi := math.Pi
	fmt.Println(pi)

	float := big.Float{}
	fmt.Println(float.SetPrec(5))
}
