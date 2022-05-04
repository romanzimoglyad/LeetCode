package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(in, 3)
	fmt.Println(in)
}

func rotate(nums []int, k int) {
	newArr := make([]int, len(nums))
	for i := range nums {
		t := (i + k) % len(nums)
		newArr[t] = nums[i]
	}
	for i := range newArr {
		nums[i] = newArr[i]
	}

}
