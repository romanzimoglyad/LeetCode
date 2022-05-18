package main

import "fmt"

func main() {
	tn6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	tn7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tn5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	tn3 := &TreeNode{
		Val:   3,
		Left:  tn6,
		Right: tn7,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  tn3,
		Right: nil,
	}
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn4,
		Right: tn5,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: nil,
	}
	// 4,2,5,1,3
	res := maxDepth(tn1)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
