package main

import "fmt"

func main() {
	d := 5
	fmt.Println(d / 2)

	t := []int{}

	fmt.Println(empty(t))

}

func empty(sw []int) bool {
	return len(sw) == 0
}
