package main

import "fmt"

func main() {
	r2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	left11 := &TreeNode{
		Val:   11,
		Left:  nil,
		Right: r2,
	}
	left4 := &TreeNode{
		Val:   4,
		Left:  left11,
		Right: nil,
	}
	right8 := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   5,
		Left:  left4,
		Right: right8,
	}
	sum := hasPathSum(root, 9)
	fmt.Println(sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
