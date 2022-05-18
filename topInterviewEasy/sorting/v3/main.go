package main

import "fmt"

func main() {
	nums1 := []int{0, 0}
	nums2 := []int{1, 8}
	merge(nums1, 0, nums2, 2)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	resP := m + n - 1
	for i := resP; i >= 0; i-- {
		if p2 < 0 {
			return
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}
