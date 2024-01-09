package main

import "fmt"

// Взятие адреса элемента map.
func main() {
	m := make(map[int]int)
	m[1] = 10
	a := &m[1]
	fmt.Println(m[1], *a)
}
