package main

import "fmt"

func main() {
	/*resp1 := fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1})
	fmt.Println(resp1)*/
	resp2 := fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	fmt.Println(resp2)
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m12 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			m12[nums1[i]+nums2[j]]++
		}
	}
	m34 := make(map[int]int)
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			m34[nums3[i]+nums4[j]]++
		}
	}
	resp := 0
	for i, vali := range m12 {
		for j, valj := range m34 {
			if i+j == 0 {
				resp += vali * valj
			}
		}
	}
	return resp
}
