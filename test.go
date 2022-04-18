package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 2
	fmt.Println(m[1])
	t := m[1]
	fmt.Println(t)
	fmt.Println(&t)
	fmt.Println(*&t)
}
