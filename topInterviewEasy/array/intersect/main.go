package main

import "fmt"

func main() {
	t := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(t)
}
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for i := range nums1 {
		m[nums1[i]]++
	}
	for j := range nums2 {
		if m[nums2[j]] != 0 {
			res = append(res, nums2[j])
			m[nums2[j]]--
		}

	}
	return res
}
