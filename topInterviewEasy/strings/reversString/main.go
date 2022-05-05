package main

import "fmt"

func main() {
	in := []byte("ABCDE")
	reverseString(in)
	fmt.Println(string(in))
}

func reverseString(s []byte) {
	fmt.Println(len(s) / 2)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
