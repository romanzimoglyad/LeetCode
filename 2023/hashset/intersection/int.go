package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 5, 6, 7}
	nums2 := []int{4, 5, 5, 6, 7, 8}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for i := range nums1 {
		m[nums1[i]] = struct{}{}
	}
	res := []int{}

	for j := range nums2 {
		if _, ok := m[nums2[j]]; ok {

			res = append(res, nums2[j])
			delete(m, nums2[j])
		}

	}
	return res
}
