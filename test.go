package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3}

	t(test...)

	t := 1
	var a *int
	a = &t
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(&a)
	fmt.Println(*&a)

}

func t(a ...int) {
	for i := range a {
		fmt.Println(a[i])
	}
}
