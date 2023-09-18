package main

import (
	"fmt"
)

func main() {
	str := []int{6, 2, 6, 5, 1, 2}

	res := removeElement(str, 2)
	fmt.Println(res)
}
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
