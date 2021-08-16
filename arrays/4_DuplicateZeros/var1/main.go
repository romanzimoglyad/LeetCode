package main

import (
	"fmt"
	funcs2 "roman_study/arrays/4_DuplicateZeros/funcs"
)

func main() {
	v := []int{1, 0, 2, 3, 0, 4, 5, 0}
	funcs2.DuplicateZeros2(v)
	fmt.Println(v)
}
