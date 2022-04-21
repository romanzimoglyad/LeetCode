package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := evalRPN([]string{"4", "3", "-"})
	fmt.Println(res)
}

func evalRPN(tokens []string) int {
	operands := map[string]struct{}{
		"*": {},
		"+": {},
		"-": {},
		"/": {}}
	m := []int{}
	for i := range tokens {
		if _, ok := operands[tokens[i]]; ok {
			switch tokens[i] {
			case "*":
				res := m[len(m)-1] * m[len(m)-2]
				m = m[0 : len(m)-2]
				m = append(m, res)
			case "+":
				res := m[len(m)-1] + m[len(m)-2]
				m = m[0 : len(m)-2]
				m = append(m, res)
			case "-":
				res := m[len(m)-2] - m[len(m)-1]
				m = m[0 : len(m)-2]
				m = append(m, res)
			case "/":
				res := m[len(m)-2] / m[len(m)-1]
				m = m[0 : len(m)-2]
				m = append(m, res)
			}

		} else {
			val, _ := strconv.Atoi(tokens[i])
			m = append(m, val)
		}
	}
	return m[0]
}
