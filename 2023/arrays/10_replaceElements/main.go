package main

import "fmt"

func main() {
	nums1 := []int{17, 18, 5, 4, 6, 1}

	res := replaceElements(nums1)
	fmt.Println(res)
}
func replaceElements(arr []int) []int {
	max := maxAfterPos(0, arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] < max {
			arr[i] = max
		} else {
			if i == len(arr)-1 {
				arr[i] = -1
				return arr
			}
			max = maxAfterPos(i+1, arr)
			arr[i] = max
		}
	}
	return arr
}

func maxAfterPos(i int, arr []int) int {
	max := arr[i]
	for i := i; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
