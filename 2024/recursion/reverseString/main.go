package main

import "fmt"

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	if len(s) == 1 || len(s) == 0 {
		return
	}

	s[0], s[len(s)-1] = s[len(s)-1], s[0]

	reverseString(s[1 : len(s)-1])
	return
}
