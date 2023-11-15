package main

import "fmt"

func f(s []string, level int) {
	a := []int{1, 2, 3}
	b := a
	b[0] = 111
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	var t []int
	fmt.Println(t)
	t = append(t, 1)
	fmt.Println(t)
}
