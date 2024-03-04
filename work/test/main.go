// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
)

func main() {
	a := []int{1}
	test(a)
	fmt.Println(a)
}

func test(in []int) {
	in = append(in, 1)
}
