package main

import "fmt"

func main() {
	t := climbStairs()
	fmt.Println(t)
}

func climbStairs(n int) int {
	m := make(map[int]int)
	return help(n, m)
}

func help(rem int, m map[int]int) int {
	if rem == 0 {
		return 1
	}
	if rem == 1 {
		return 1
	}
	var var1, var2 int
	if val, ok := m[rem-1]; ok {
		var1 = val
	} else {
		var1 = help(rem-1, m)
	}
	if val, ok := m[rem-2]; ok {
		var2 = val
	} else {
		var2 = help(rem-2, m)
	}

	m[rem] = var1 + var2
	return var1 + var2
}
