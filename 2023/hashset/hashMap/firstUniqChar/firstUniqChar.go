package main

import "fmt"

func main() {
	arr := "loveleetcode"
	res := firstUniqChar(arr)
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for _, val := range s {
		m[val]++
	}

	for i, val := range s {
		if m[val] == 1 {
			return i
		}
	}
	return -1
}
