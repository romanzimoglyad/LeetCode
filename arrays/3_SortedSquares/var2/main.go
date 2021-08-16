package main

import (
	"fmt"
	"roman_study/arrays/3_SortedSquares/funcs"
)

func main() {
	v := []int{-7, -3, 2, 3, 11}
	t := funcs.SortedSquares2(v)
	fmt.Println(t)
}
