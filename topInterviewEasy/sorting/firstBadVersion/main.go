package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func main() {
	fmt.Println(firstBadVersion(2))
}

func firstBadVersion(n int) int {

	return helper(0, n)
}
func helper(start, end int) int {
	if start == end {
		if isBadVersion(start) {
			return start
		}
		if isBadVersion(end) {
			return end
		}
	}
	mid := start + (end-start+1)/2

	if !isBadVersion(mid) {
		return helper(mid, end)
	} else {
		return helper(start+1, mid)
	}
}
func isBadVersion(version int) bool {
	return version >= 0
}
