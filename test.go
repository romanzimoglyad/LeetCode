package main

import (
	"fmt"
)

type Person struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

const format = "2006-01-02 15:04:05 +0000 UTC"

func main() {
	var m map[int]int

	fmt.Println(m[1])
}

func add(arr []int) {
	arr = append(arr, 2)
	fmt.Println(len(arr))
	fmt.Println(arr)
}
