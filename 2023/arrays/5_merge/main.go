package main

import "fmt"

func main() {
	nums1 := []int{1, 5, 7, 8, 9, 0, 0}
	nums2 := []int{2, 4}
	merge(nums1, 3, nums2, 2)
	fmt.Println(nums1)
	t := max(1, 5)
	fmt.Println(t)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	resp := make([]int, 0, n)
	left1 := 0
	left2 := 0
	for left1 <= m-1 && left2 <= n-1 {

		if nums1[left1] <= nums2[left2] {
			resp = append(resp, nums1[left1])
			left1++
		} else {
			resp = append(resp, nums2[left2])
			left2++
		}
	}

	if left2 == n {
		resp = append(resp, nums1[left1:]...)

	} else {
		resp = append(resp, nums2[left2:]...)

	}
	copy(nums1, resp)
}
