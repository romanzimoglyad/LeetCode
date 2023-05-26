package main

import "fmt"

func main() {
	nums1 := []int{3, 1, 7, 11}

	res := checkIfExist(nums1)
	fmt.Println(res)
}

func checkIfExist(arr []int) bool {
	m := make(map[int]struct{})
	for i := range arr {
		if _, ok := m[arr[i]]; ok {
			return true
		}

		if arr[i]%2 == 0 {
			m[arr[i]/2] = struct{}{}
		}
		m[arr[i]*2] = struct{}{}
	}
	return false
}
