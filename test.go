package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1
	for 100*n*n >= int(math.Pow(2, float64(n))) {
		n++
	}

	fmt.Println("The smallest n where 100n^2 < 2^n is approximately:", n)
}
