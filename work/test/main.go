// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
	"strings"
)

func main() {
	resp := decode_str("3[a2[c]]")
	fmt.Println(resp)
}

func decode_str(s string) string {
	res := ""
	curr_num := 0
	depth := 0
	buf := "" // contains '[]' contents

	for _, ch := range s {
		if depth == 0 && ch >= '0' && ch <= '9' {
			curr_num = curr_num*10 + int(ch-'0')
			continue
		}

		if ch == '[' {
			depth++
			if depth == 1 { // don't add '[' in the 0 nesting level
				continue
			}
		}

		if ch == ']' {
			depth--
			if depth == 0 {
				res += strings.Repeat(decode_str(buf), curr_num)
				buf = ""
				curr_num = 0
				continue
			}
		}

		if depth == 0 {
			res += string(ch)
		} else {
			buf += string(ch)
		}
	}

	return res
}
