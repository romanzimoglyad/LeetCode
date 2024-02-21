// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
)

func main() {
	var test float64
	test = 164000000
	g := fmt.Sprintf("%.2f", test)
	fmt.Println(g)
}
