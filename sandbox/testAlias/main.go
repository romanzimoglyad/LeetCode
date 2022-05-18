package main

import (
	"fmt"
	"strings"
)

type test struct {
	t string
	v rune
}
type alias test

func (t test) Say() {
	fmt.Println("test")
}
func (t alias) say1() {
	fmt.Println("alias")
}
func main() {
	var b strings.Builder
	b.Grow(32)
	for i, p := range []int{2, 3, 5, 7, 11, 13} {
		fmt.Fprintf(&b, "%d:%d, ", i+1, p)
	}
	s := b.String() // нет копирования

	s = s[:b.Len()-2] // нет копирования
	// (удаляет завершающий ", ")
	fmt.Println(s)

	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c %T %v \n ", i, c, c, c)
	}
}
