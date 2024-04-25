package main

import "fmt"

func main() {

	arr := []int{1, 4, 6, 3, 5, 7}
	arr = []int{6, 5, 4, 3, 2, 1}
	fmt.Println(intersections(arr))
}

var intCount int

func intersections(arr []int) int {
	intCount = 0
	intRec(arr)
	return intCount
}

func intRec(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := intRec(arr[:mid])
	right := intRec(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	lenLeft := len(left)
	lenRight := len(right)
	res := make([]int, 0, lenLeft+lenRight)
	i, j := 0, 0
	for i < lenLeft && j < lenRight {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			intCount = intCount + lenLeft - i
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
