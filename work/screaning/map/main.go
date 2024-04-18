package main

import "fmt"

// Взятие адреса элемента map.
func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	y := 1
	v := m["a"]
	fmt.Println(v)
	t := *y
	fmt.Println(t)
	for a, b := range m {
		fmt.Println(a, b)
	}

}
