package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	sl1 := make([]*int, 0, 5)
	for _, elem := range sl {
		sl1 = append(sl1, &elem)

	}
	for _, elem := range sl1 {
		fmt.Println(*elem)

	}
}
