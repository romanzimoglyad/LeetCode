package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
}

func isRot(s1, s2 string) bool {
	newS2 := s2 + s2

	return isSubstring(newS2, s1)

}

func isSubstring(s1, s2 string) bool {

	return strings.Contains(s1, s2)
}
