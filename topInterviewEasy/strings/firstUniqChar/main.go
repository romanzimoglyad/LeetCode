package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	m := make(map[byte]struct{})
	b := []byte(s)
	for i := range b {
		val := b[i]

		if _, ok := m[val]; !ok && strings.LastIndex(s, string(val)) == i {
			return i
		}
		m[val] = struct{}{}
	}
	return -1
}
