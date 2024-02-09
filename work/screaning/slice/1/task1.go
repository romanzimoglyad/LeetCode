package main

import "fmt"

func main() {
	a()
}

//func a() {
//	x := []int{}     // len=0 cap =0
//	x = Append(x, 0) // x: len=1 cap = 1
//	x = Append(x, 1) // x: len=2 cap = 2
//	x = Append(x, 2) // x: len= 3 cap = 4
//
//	y := Append(x, 3)
//	fmt.Println(x, y)    // x: len=3 cap = 4  y: len=4 cap =4
//	z := Append(x, 4)    // x: len=3 cap = 4  y: len=4 cap =4 z: len=4 cap =4
//	fmt.Println(x, y, z) // len= cap =
//}


func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}


/*
func main() {
	a()
}
func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(x, y, z)
}


*/


func a() {
	x := make([]int, 0, 3) //x len 0  cap 3
	x = append(x, 0) // len 3 cap 4
	x = append(x, 1)// len 4 cap 4
	x = append(x, 2)// len 5 cap  8

	y := append(x, 3) // len 5 cap  8
	z := append(x, 4)
	fmt.Println(y, z)
}
