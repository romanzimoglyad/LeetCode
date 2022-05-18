package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("2147483648"))
}
func myAtoi(s string) int {

	s1 := strings.Trim(s, " ")
	if len(s1) == 0 {
		return 0
	}

	neg := false
	if s1[0] == '-' {
		neg = true
		s1 = s1[1:]
	} else if s1[0] == '+' {
		s1 = s1[1:]
	}
	resB := []byte{}

	for i := range s1 {
		if s1[i] >= '0' && s1[i] <= '9' {
			resB = append(resB, s1[i])
		} else {
			break
		}
	}
	if len(resB) == 0 {
		return 0
	}
	fmt.Println(math.MaxInt32 / 10)
	var res int
	for _, elem := range resB {
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && int(elem-'0') > math.MaxInt32%10) {
			if neg {
				return math.MinInt32
			} else {
				res = math.MaxInt32
			}
			break
		}
		res = res*10 + int(elem-'0')
	}

	if neg {
		res = res * -1
	}
	return res
}
