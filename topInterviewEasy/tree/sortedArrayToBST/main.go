package main

import "fmt"

func main() {
	in := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(in)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums) / 2

	resp := &TreeNode{Val: nums[mid]}
	resp.Left = sortedArrayToBST(nums[0:mid])
	resp.Right = sortedArrayToBST(nums[mid+1:])
	return resp
}
