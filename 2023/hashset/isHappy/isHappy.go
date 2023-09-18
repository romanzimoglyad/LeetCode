package main

import "fmt"

func main() {
	n := 2
	res := isHappy(n)
	fmt.Println(res)
}
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for {
		res := getSquare(n)
		if res == 1 {
			return true
		}
		if _, ok := m[res]; ok {
			return false
		}
		m[res] = struct{}{}
		n = res
	}
	return false
}

func getSquare(n int) int {
	res := 0
	for n != 0 {
		v := n % 10
		sq := v * v
		res += sq
		n /= 10
	}
	return res
}
