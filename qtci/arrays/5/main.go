package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(isPalindromPermutatuin("aaabbb"))
	//fmt.Println(isPalindromPermutatuin("aaa"))
	//fmt.Println(isPalindromPermutatuin("ab"))
	fmt.Println(isPalindromPermutatuin("Tact Coa"))
	fmt.Println(isPalindromPermutatuin("abab"))
	fmt.Println(isPalindromPermutatuin("cbbb"))
}
func isPalindromPermutatuin(in string) bool {
	m := make(map[rune]int)
	s := strings.ReplaceAll(in, " ", "")
	s = strings.ToLower(s)
	for _, letter := range s {
		m[letter]++
	}
	oddFlag := false
	for _, v := range m {
		if v%2 != 0 {
			if oddFlag {
				return false
			}
			oddFlag = true
		}
	}
	return true
}
