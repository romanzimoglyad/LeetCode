package main

import "fmt"

func main() {

	fmt.Println(mSort1([]int{2, 1, 23, 88, 8888, 99, 4, 5, 8, 5, 3, 2}))
}

func mSort1(in []int) []int {
	if len(in) == 1 || len(in) == 0 {
		return in
	}
	mid := len(in) / 2
	left := mSort1(in[0:mid])
	right := mSort1(in[mid:])
	return mSort(left, right)
}

func mSort(n1 []int, n2 []int) []int {
	res := make([]int, 0, len(n1)+len(n2))
	for len(n1) > 0 || len(n2) > 0 {
		if len(n2) == 0 {
			return append(res, n1...)
		}
		if len(n1) == 0 {
			return append(res, n2...)
		}
		if n1[0] > n2[0] {
			res = append(res, n2[0])
			n2 = n2[1:]
		} else {
			res = append(res, n1[0])
			n1 = n1[1:]
		}
	}
	return res
}
