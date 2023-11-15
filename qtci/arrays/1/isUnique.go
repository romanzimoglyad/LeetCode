package main

import "fmt"

func main() {
	fmt.Println(isUnique("париа"))
	fmt.Println(isUnique2("aaa"))
}

func isUnique(s string) bool {

	m := make(map[rune]struct{})
	for i, v := range s {
		fmt.Printf("s[i] %T, %q \n", s[i], s[i])
		fmt.Printf("v %T, %q\n  ", v, v)
		fmt.Printf("a %T %v\n", 'a', 'a')
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v] = struct{}{}
		}
	}
	return true
}
func isUnique2(s string) bool {
	bm := make([]bool, 122-97)
	for _, v := range s {
		if bm[v-'a'] == true {
			return false
		} else {
			bm[v-'a'] = true
		}
	}
	return true
}
