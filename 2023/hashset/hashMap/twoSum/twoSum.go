package main

import "fmt"

func main() {
	arr := []int{3, 2, 4}
	res := twoSum(arr, 6)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for j, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{j, val}
		} else {
			m[num] = j
		}

	}
	return nil
}
