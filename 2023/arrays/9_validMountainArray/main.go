package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 1}

	res := validMountainArray(nums1)
	fmt.Println(res)
}

func validMountainArray(arr []int) bool {
	peak := false
	up := false
	if len(arr) < 3 {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			return false
		}
		if arr[i-1] < arr[i] && !peak {
			up = true
			continue
		}

		if arr[i-1] >= arr[i] && !peak {
			peak = true
		}
		if arr[i-1] <= arr[i] && peak {
			return false
		}
	}
	return peak && up
}
