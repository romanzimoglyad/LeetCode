package main

import "fmt"

func main() {
	/*	r2 := &TreeNode{
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
		}*/
	left15 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	r15 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	left1 := &TreeNode{
		Val:   5,
		Left:  left15,
		Right: r15,
	}
	r55 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	r5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: r55,
	}
	root := &TreeNode{
		Val:   5,
		Left:  left1,
		Right: r5,
	}

	/*	ll := &TreeNode{
		Val:   111,
		Left:  nil,
		Right: nil,
	}*/
	r := countUnivalSubtrees(root)
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	c := 0

	isUniVal(root, &c)

	return c
}

func isUniVal(root *TreeNode, count *int) bool {

	if root == nil {
		return true
	}
	parent := root.Val
	if root.Left == nil && root.Right == nil {
		*count++
		return true
	}
	left := true
	right := true
	if root.Left != nil {
		left = isUniVal(root.Left, count) && root.Left.Val == parent
	}
	if root.Right != nil {
		right = isUniVal(root.Right, count) && root.Right.Val == parent
	}
	if left && right {
		*count++
		return true
	}
	return false
}
