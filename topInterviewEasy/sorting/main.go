package main

import "fmt"

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	res := make([]int, 0, m+n)
	i := 0
	j := 0
	for i < m || j < n {
		if j == n {
			res = append(res, nums1[i:]...)
			break
		}
		if i == m {
			res = append(res, nums2[j:]...)
			break
		}
		if nums1[i] > nums2[j] {
			res = append(res, nums2[j])

			j++
		} else {
			res = append(res, nums1[i])
			i++
		}
	}
	copy(nums1, res)
}
