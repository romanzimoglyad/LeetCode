package main

import "fmt"

func main() {
	resp := numJewelsInStones("aA", "aAAdsdf")
	fmt.Println(resp)
}

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]struct{})
	for _, v := range jewels {
		m[v] = struct{}{}
	}
	resp := 0
	for _, v := range stones {
		if _, ok := m[v]; ok {
			resp++
		}
	}
	return resp
}
