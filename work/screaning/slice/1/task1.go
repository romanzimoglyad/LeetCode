package main

import "fmt"

func main() {
	a()
}

func a() {
	x := []int{}         // len=0 cap =0
	x = append(x, 0)     // x: len=1 cap = 1
	x = append(x, 1)     // x: len=2 cap = 2
	x = append(x, 2)     // x: len= 3cap = 4
	y := append(x, 3)    // x: len=3 cap = 4  y: len=4 cap =4
	z := append(x, 4)    // x: len=3 cap = 4  y: len=4 cap =4 z: len=4 cap =4
	fmt.Println(x, y, z) // len= cap =
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
