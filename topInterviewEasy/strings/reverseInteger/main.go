package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	pr := strconv.Itoa(x)
	ne := false
	if string(pr[0]) == "-" {
		ne = true
		pr = pr[1:]
	}
	str := []byte(pr)
	for i, j := 0, len(str)-1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	str1 := string(str)
	out2, err := strconv.Atoi(str1)
	if out2 >= 2147483648 {
		return 0
	}
	if ne {
		str1 = "-" + str1
	}

	out, err := strconv.Atoi(str1)
	if err != nil || out >= 2147483648 {
		return 0
	}
	return out
}
