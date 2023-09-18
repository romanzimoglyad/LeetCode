package main

import "fmt"

func main() {

	res := generate(5)
	fmt.Println(res)
}

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	prev := []int{1}
	res = append(res, prev)
	for i := 1; i < numRows; i++ {
		newArr := generateNewArr(prev)
		res = append(res, newArr)
		prev = newArr
	}
	return res
}

func generateNewArr(prevArr []int) []int {
	newArr := make([]int, 0, len(prevArr)+1)
	extendedPrevArr := make([]int, 0, len(prevArr)+2)
	extendedPrevArr = append([]int{0}, prevArr...)
	extendedPrevArr = append(extendedPrevArr, 0)
	for i := 0; i < len(extendedPrevArr)-1; i++ {
		newVal := extendedPrevArr[i] + extendedPrevArr[i+1]
		newArr = append(newArr, newVal)
	}
	return newArr
}
