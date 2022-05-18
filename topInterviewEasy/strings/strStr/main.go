package main

import "fmt"

func main() {
	fmt.Println(strStr("aabaabbbaabbbbabaaab", "abaa"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	pos := 0
	tmpPos := 0
	needlePos := 0
	for pos < len(haystack) || needlePos == len(needle) {
		if needlePos == len(needle) {
			return pos - needlePos
		}
		if haystack[pos] == needle[needlePos] {
			if tmpPos == 0 {
				tmpPos = pos + 1
			}
			pos++
			needlePos++
		} else {
			needlePos = 0
			if tmpPos != 0 {
				pos = tmpPos
				tmpPos = 0
			} else {
				pos++
			}

		}
	}
	return -1
}
