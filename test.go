package main

import (
	"fmt"
)

func main() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}
func desc(s []int) {
	fmt.Printf("len(%d), cap(%d),arr:%v\n", len(s), cap(s), s)
}
