package main

import "fmt"

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, len(nums1))
	i2 := 0
	i1 := 0
	for i2 < n && i1 < m {
		v2 := nums2[i2]
		v1 := nums1[i1]
		if v2 > v1 {
			res = append(res, nums1[i1])
			i1++
		} else {
			res = append(res, v2)
			i2++
		}
	}
	if i1 != m {
		for i1 := i1; i1 < m; i1++ {
			res = append(res, nums1[i1])
		}
	}
	if i2 != n {
		for i2 := i2; i2 < n; i2++ {
			res = append(res, nums2[i2])
		}
	}
	copy(nums1, res)
}
