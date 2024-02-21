package main

import "fmt"

func main() {
	r := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := l / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
