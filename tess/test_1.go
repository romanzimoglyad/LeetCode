package main

import "fmt"

func main() {
	arr := make([]int, 10, 10)
	fmt.Println(len(arr), cap(arr))
	arr[5] = 1
	fmt.Println(len(arr), cap(arr))
}
