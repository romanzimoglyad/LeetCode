package main

import "fmt"

func main() {
	slice1 := make([]int, 5)
	test1(slice1)
	fmt.Println(slice1)
	slice1 = make([]int, 5)
	test2(slice1)
	fmt.Println(slice1)
}
func test1(slice []int) {
	for i := range slice {
		slice[i] = slice[i] + 10
	}
}

func test2(slice []int) {
	for i := 0; i < 5; i++ {
		slice = append(slice, i*2)
	}
}
