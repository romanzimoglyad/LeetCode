package main

import (
	"fmt"
)

func main() {
	res := groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"})
	fmt.Println(res)
	/*
		resp := shift("abc")
		fmt.Println(resp)*/
}

func groupStrings(strings []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strings {
		shifts := shift(v)
		f := false
		for i := 0; i < len(shifts); i++ {
			if _, ok := m[shifts[i]]; ok {
				m[shifts[i]] = append(m[shifts[i]], v)
				f = true
				break
			}
		}
		if !f {
			m[v] = append(m[v], v)
		}
	}
	var r [][]string
	for _, v := range m {
		r = append(r, v)
	}

	return r
}

func shift(s string) [26]string {
	resp := [26]string{}
	for i := 0; i < 26; i++ {
		var newVar []rune
		for j := 0; j < len(s); j++ {
			newChar := 97 + (s[j]+1)%97%26
			newVar = append(newVar, rune(newChar))
		}
		resp[i] = string(newVar)
		s = string(newVar)
	}
	return resp
}
