package main

import "fmt"

func main() {
	test := []int{1, 10, 5, 7, 11, 555, 1, 5, 110, 10, 10, 5}
	fmt.Println(mergeSort(test))
}

func mergeSort(a []int) []int {
	l := len(a)
	if l == 1 {
		return a
	}
	mid := l / 2
	arr1 := mergeSort(a[:mid])
	arr2 := mergeSort(a[mid:])
	return merge(arr1, arr2)
}

// 1,3
// 2
func merge(a, b []int) []int {
	l1 := len(a)
	l2 := len(b)
	res := make([]int, 0, l1+l2)
	i := 0
	j := 0
	for i < l1 && j < l2 {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, b[j:]...)
	res = append(res, a[i:]...)

	return res
}
