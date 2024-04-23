package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

const format = "2006-01-02 15:04:05 +0000 UTC"

func main() {
	t := "2023-11-28 16:00:00 +0000 UTC"
	dateAccrual, err := time.Parse(format, t)
	fmt.Println(err)
	fmt.Println(dateAccrual)
}

func add(arr []int) {
	arr = append(arr, 2)
	fmt.Println(len(arr))
	fmt.Println(arr)
}
