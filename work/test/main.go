// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "Победа (рейс выполняет туруру)"
	test = strings.ReplaceAll(test, " (", ", ")
	test = strings.ReplaceAll(test, ")", "")
	fmt.Println(test)
}
