package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 2}

	n := validMountainArray(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	down := false
	peak := false
	hasUp := false
	i := 1
	for i < len(arr) {
		prev := arr[i-1]
		curr := arr[i]
		if prev == curr {
			return false
		}
		if prev < curr && !down {
			hasUp = true
			i++
			continue
		}
		if prev < curr && down {
			return false
		}
		if prev > curr && down {
			i++
			continue
		}
		if prev > curr && !down {
			if !peak {
				peak = true
				down = true
			} else {
				return false
			}

		}
	}

	return peak && hasUp
}
