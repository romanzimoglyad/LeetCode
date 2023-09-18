package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "good example"

	res := reverseWords(s)
	fmt.Println(res)
}
func reverseWords(s string) string {
	sss := strings.Split(s, " ")
	res := make([]string, 0, len(sss))
	for _, ss := range sss {
		res = append(res, reverseString(ss))
	}
	return strings.Join(res, " ")
}

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
