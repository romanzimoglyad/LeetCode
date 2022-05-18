package main

import "fmt"

func main() {
	fmt.Println(isAnagram("aacc", "ccac"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	b1 := []byte(s)
	b2 := []byte(t)
	m := make(map[byte]*int)
	for i := 0; i < len(b1); i++ {
		if val, ok := m[b1[i]]; !ok {
			t := 1
			m[b1[i]] = &t

		} else {
			*val++
		}

	}
	for i := 0; i < len(b2); i++ {
		if count, ok := m[b2[i]]; !ok || *count == 0 {
			return false
		} else {
			*count--
		}

	}
	return true
}
