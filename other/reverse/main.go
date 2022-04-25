package main

import "fmt"

func main() {
	d := 5
	fmt.Println(d / 2)

	t := []int{1, 2, 3, 4, 5}
	reverse(t)

	fmt.Println(t)

}

func reverse(sw []int) {
	for i := 0; i < len(sw)/2; i++ {
		sw[i], sw[len(sw)-1-i] = sw[len(sw)-i-1], sw[i]
	}
}
