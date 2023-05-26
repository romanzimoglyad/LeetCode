package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i] - i
	}

	minDiff := b[0]
	maxDiff := b[0]
	for i := 1; i < n; i++ {
		diff := b[i] - b[i-1]
		if diff < minDiff {
			minDiff = diff
		}
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	if maxDiff-minDiff > 2 {
		fmt.Println("x")
	} else if maxDiff-minDiff == 2 {
		for i := 0; i < n-1; i++ {
			if b[i+1]-b[i] == 2 {
				fmt.Printf("%d ", (a[i]+a[i+1])/2)
			} else {
				fmt.Printf("%d ", a[i])
			}
		}
		fmt.Println(a[n-1])
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", a[i])
		}
		fmt.Println()
	}
}
