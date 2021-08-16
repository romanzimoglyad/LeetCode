package main

import "fmt"

func main() {
	nums1 := []int{17, 18, 5, 4, 6, 1}

	n := replaceElements(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

func replaceElements(arr []int) []int {
	max := -1

	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = max
		if tmp > max {
			max = tmp
		}

	}

	return arr
}
