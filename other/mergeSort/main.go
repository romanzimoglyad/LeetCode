package main

import "fmt"

func main() {
	data := []int{9, 2, 7, 5, 8, 1, 6, 3, 9, 33, 0}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))
}

func MergeSort(data []int) []int {
	if len(data) == 1 {
		return data
	}
	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)

		} else if len(right) == 0 {
			return append(res, left...)

		} else if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]

		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}
