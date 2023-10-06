package main

import "fmt"

func main() {
	nums := []int{1}
	k := 1
	res := topKFrequent(nums, k)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	maxFreq := 0
	for _, v := range nums {
		m[v]++
		if m[v] > maxFreq {
			maxFreq = m[v]
		}
	}
	r := make([][]int, maxFreq+1)
	for k, v := range m {
		r[v] = append(r[v], k)
	}
	res := make([]int, 0, k)

	for i := len(r) - 1; i >= 0 && k > 0; i-- {
		if len(r[i]) != 0 {
			for j := 0; j < len(r[i]) && k > 0; j++ {
				res = append(res, r[i][j])
				k--
			}
		}
	}

	return res
}
