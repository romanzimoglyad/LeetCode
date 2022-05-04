package main

import "fmt"

func main() {
	in := []int{2, 7, 11, 15}
	i1 := []int{3, 3}
	fmt.Println(twoSum(in, 9))
	fmt.Println(twoSum(i1, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if j, ok := m[nums[i]]; ok {
			return []int{i, j}
		}
		m[target-nums[i]] = i
	}
	return nil
}
