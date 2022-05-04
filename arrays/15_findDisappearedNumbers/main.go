package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4, 3, 2, 1}

	n := findDisappearedNumbers2(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

func findDisappearedNumbers(nums []int) []int {
	t := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		t[nums[i]] = 1
	}
	res := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := t[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}
func findDisappearedNumbers1(nums []int) []int {
	appeareances_arr := make([]bool, len(nums))
	for _, n := range nums {
		if n > 0 && n <= len(nums) {
			appeareances_arr[n-1] = true
		}
	}
	missing_numbers := []int{}
	for idx, appeared := range appeareances_arr {
		if appeared {
			continue
		}
		missing_numbers = append(missing_numbers, idx+1)
	}
	return missing_numbers
}
func findDisappearedNumbers2(nums []int) []int {
	// fmt.Println(nums)
	for i := 0; i < len(nums); {
		pos := nums[i] - 1
		if i != pos {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}

	// fmt.Println(nums)

	miss := make([]int, 0, 0)

	for i, v := range nums {
		if v != i+1 {
			miss = append(miss, i+1)
		}
	}
	// fmt.Println(miss)
	return miss
}
