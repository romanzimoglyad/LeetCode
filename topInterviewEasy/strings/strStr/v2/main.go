package main

import "fmt"

func main() {
	fmt.Println(strStr("abc", "bc"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
