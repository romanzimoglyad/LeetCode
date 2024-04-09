package main

import "fmt"

func main() {
	r := tribonacci(25)
	fmt.Println(r)
}

var m = map[int]int{}

func tribonacci(n int) int {
	m = make(map[int]int)
	return dp(n)
}

func dp(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 || i == 2 {
		return 1
	}
	if _, ok := m[i]; !ok {
		m[i] = dp(i-1) + dp(i-2) + dp(i-3)
	}

	return m[i]

}
