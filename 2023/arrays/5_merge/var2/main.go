package main

import "fmt"

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	m--
	n--
	for k := m + n + 1; k >= 0; k-- {
		if n < 0 {
			break
		}

		if m >= 0 && nums1[m] > nums2[n] {
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[n]
			n--
		}
	}

}
