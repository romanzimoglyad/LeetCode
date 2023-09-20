package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 3, 4, 5, 6, 7, 8}
	nums2 := []int{1, 2, 3, 4, 3, 3, 4}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for i := range nums1 {
		m[nums1[i]]++
	}
	for i := range nums2 {
		if val, ok := m[nums2[i]]; ok && val != 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}
	return res
}
