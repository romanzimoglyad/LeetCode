package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	var s1 []string
	for _, el := range s {
		if unicode.IsLetter(el) || unicode.IsDigit(el) {
			s1 = append(s1, strings.ToUpper(string(el)))
		}
	}
	for i, j := 0, len(s1)-1; i < j; {
		if s1[i] != s1[j] {
			return false
		}
		i++
		j--
	}
	return true
}
