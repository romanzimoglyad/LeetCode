package main

import "fmt"

func main() {
	data := []int{9, 2}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))

	/*	data1 := []int{1, 5, 8}
		data2 := []int{2, 6, 10}
		fmt.Println(Merge1(data1, data2))*/
}

func MergeSort(data []int) []int {
	if len(data) == 0 || len(data) == 1 {
		return data
	}
	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return merge1(left, right)
}

func merge22(left []int, right []int) []int {
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
func merge1(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
