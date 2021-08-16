package main

import (
	"fmt"
)

func main() {
	v := []int{555, 901, 482, 1771}
	t := findNumbers(v)
	fmt.Println(t)
}

func findNumbers(nums []int) int {
	res := 0
	for i := range nums {
		str := iterativeDigitsCount(nums[i])
		if str%2 == 0 {
			res++
		}
	}
	return res
}

func iterativeDigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count

}
