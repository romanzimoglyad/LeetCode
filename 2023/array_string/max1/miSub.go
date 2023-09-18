package main

import "fmt"

func main() {
	str := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(str, 3)
	fmt.Println(str)
}
func rotate(nums []int, k int) {
	l := len(nums)
	newPos := k % l
	if newPos == 0 {
		return
	}

	for counter := 0; counter < l-newPos; counter++ {
		mem := 0
		inMem := false
		for i := 0; i < l-1; i++ {
			if !inMem {
				inMem = true
				mem = nums[i]
			}
			nums[i] = nums[i+1]
		}
		nums[l-1] = mem
	}
	return
}
