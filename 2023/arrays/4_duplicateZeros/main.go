package main

import "fmt"

func main() {
	testArr := []int{1, 0, 2, 3, 0, 0, 4}
	duplicateZeros(testArr)
	fmt.Println(testArr)
}
func duplicateZeros(arr []int) {
	l := len(arr)
	res := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if arr[i] != 0 {
			res = append(res, arr[i])
		} else {
			res = append(res, arr[i])
			res = append(res, 0)
		}
	}
	copy(arr, res)
}
