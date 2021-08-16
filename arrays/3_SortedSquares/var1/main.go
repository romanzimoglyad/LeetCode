package main

import (
	"fmt"
	"roman_study/arrays/3_SortedSquares/funcs"
)

func main() {
	v := []int{-4, -1, 0, 3, 10}
	t := funcs.SortedSquares1(v)
	fmt.Println(t)
}
