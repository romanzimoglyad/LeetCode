package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a good   example"

	res := reverseWords(s)
	fmt.Println(res)
}
func reverseWords(s string) string {
	r := reverseArr(strings.Split(strings.TrimSpace(s), " "))
	return strings.Join(r, " ")
}

func reverseArr(str []string) []string {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	res := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] != "" {
			res = append(res, str[i])
		}
	}

	return res
}
