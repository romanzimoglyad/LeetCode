package main

import "fmt"

func main() {
	arg := "paper"
	res := isIsomorphic(arg, "title")
	fmt.Println(res)

	// paper title
}

func isIsomorphic(s string, t string) bool {
	m := make(map[string]string)
	var resStr string
	for i, val := range s {
		if v, ok := m[string(val)]; ok {
			resStr = resStr + v
		} else {
			for _, v := range m {
				if v == string(t[i]) {
					return false
				}
			}

			resStr = resStr + string(t[i])
			m[string(val)] = string(t[i])
		}
		if resStr != t[0:i+1] {
			return false
		}
	}

	return true
}
