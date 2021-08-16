package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 8, 11}

	n := checkIfExist(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	for i := range arr {
		if _, ok := m[arr[i]/2]; ok && arr[i]%2 == 0 {

			return true

		}
		if _, ok := m[arr[i]*2]; ok {

			return true

		}
		m[arr[i]] = true

	}
	return false
}
