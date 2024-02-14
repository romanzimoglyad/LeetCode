package main

import "fmt"

func main() {
	resp := fib(6)
	fmt.Println(resp)
}

func fib(n int) int {
	m := make(map[int]int)
	return helper(n, m)
}

func helper(n int, m map[int]int) int {
	if val, ok := m[n]; ok {
		return val
	}
	if n < 2 {
		return n
	}
	resp := helper(n-1, m) + helper(n-2, m)
	m[n] = resp
	return resp
}
