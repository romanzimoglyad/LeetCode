package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	res := containsNearbyDuplicate(nums, 2)
	fmt.Println(res)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)
	for k, v := range nums {
		m[v] = append(m[v], k)
	}

	for _, arr := range m {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1]-arr[j] <= k {
				return true
			}
		}
	}
	return false
}
