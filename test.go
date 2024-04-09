package main

import "fmt"

type Person struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	testArr := make([]int, 0, 5)
	testArr = append(testArr, 1)
	fmt.Println(len(testArr))
	add(testArr)
	fmt.Println(len(testArr))
	fmt.Println(testArr)
}

func add(arr []int) {
	arr = append(arr, 2)
	fmt.Println(len(arr))
	fmt.Println(arr)
}
