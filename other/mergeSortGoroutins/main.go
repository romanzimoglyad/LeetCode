package main

import "fmt"

func main() {
	data := []int{6, 2, 9, 1}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))
}

func MergeSort(data []int) []int {
	if len(data) == 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int
	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()

	right := MergeSort(data[mid:])
	<-done
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
