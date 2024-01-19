package main

import "fmt"

// Взятие адреса элемента map.
func main() {
	m := make(map[int]int)
	defer (fmt.Println(1))
	defer (fmt.Println(2))
}
