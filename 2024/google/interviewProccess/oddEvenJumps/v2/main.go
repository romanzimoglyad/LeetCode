package main

import "sort"

func oddEvenJumps(arr []int) int {
	n := len(arr)
	oddJump := make([]bool, n)
	evenJump := make([]bool, n)
	oddJump[n-1], evenJump[n-1] = true, true // Last index is always reachable

	nextHigher := make([]int, n) // Next higher element's index
	nextLower := make([]int, n)  // Next lower element's index

	// Building nextHigher and nextLower arrays
	stack := make([]int, 0)
	for _, i := range sortT(arr) {
		for len(stack) > 0 && stack[len(stack)-1] < i {
			nextHigher[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = stack[:0] // Clear the stack
	for _, i := range sortT(reverse(arr)) {
		for len(stack) > 0 && stack[len(stack)-1] < i {
			nextLower[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Dynamic programming to check which indices are reachable
	for i := n - 2; i >= 0; i-- {
		if nextHigher[i] > 0 {
			oddJump[i] = evenJump[nextHigher[i]]
		}
		if nextLower[i] > 0 {
			evenJump[i] = oddJump[nextLower[i]]
		}
	}

	// Counting good starting indices
	count := 0
	for _, jump := range oddJump {
		if jump {
			count++
		}
	}
	return count
}

// Helper function to get sorted indices
func sortT(arr []int) []int {
	indexes := make([]int, len(arr))
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return arr[indexes[i]] < arr[indexes[j]]
	})
	return indexes
}

// Helper function to get sorted indices in reverse order
func reverse(arr []int) []int {
	indexes := sortT(arr)
	n := len(indexes)
	for i := 0; i < n/2; i++ {
		indexes[i], indexes[n-i-1] = indexes[n-i-1], indexes[i]
	}
	return indexes
}

func main() {
	arr := []int{2, 3, 1, 1, 4}
	println(oddEvenJumps(arr)) // Output: 2
}
