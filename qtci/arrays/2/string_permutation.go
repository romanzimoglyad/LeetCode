package main

import (
	"fmt"
)

func main() {
	fmt.Println(stringPerm("abcd", "bcda"))
	fmt.Println(stringPerm("abcda", "bcdaz"))
}

func stringPerm(a, b string) bool {

	if len(a) != len(b) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}
