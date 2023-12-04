package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := comp("aaabbbdddasd")
	fmt.Println(res)
}

func comp(s string) string {
	if len(s) == 0 {
		return s
	}
	res := strings.Builder{}
	s = strings.ToLower(s)
	var ll rune
	lc := 0

	for _, v := range s {
		if v == ll {
			lc++
			continue
		}
		if ll != 0 {
			res.WriteRune(ll)
			res.WriteString(strconv.Itoa(lc))
		}

		lc = 1
		ll = v
	}
	res.WriteRune(ll)
	res.WriteString(strconv.Itoa(lc))

	return res.String()

}
