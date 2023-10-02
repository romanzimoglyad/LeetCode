package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {

	max := 0
	i := 0
	m := make(map[uint8]int)
	for j := 0; j < len(s); j++ {
		if val, ok := m[s[j]]; ok {
			i = maxim(val, i)
		}
		m[s[j]] = j + 1
		max = maxim(max, j-i+1)

	}
	return max
}
func maxim(i, j int) int {
	if i > j {
		return i
	}
	return j

}
