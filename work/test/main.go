// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
)

type Parent struct{}

func (c *Parent) Print() {
	fmt.Println("parent")
}

type Child struct {
	Parent
}

func (p *Child) Print() {
	fmt.Println("child")
}

func main() {
	var x Child
	x.Parent.Print()
}
