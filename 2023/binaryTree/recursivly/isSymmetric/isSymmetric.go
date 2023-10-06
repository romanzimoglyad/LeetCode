package main

import "fmt"

func main() {

	a4 := &TreeNode{
		Val:   22,
		Left:  nil,
		Right: nil,
	}
	a5 := &TreeNode{
		Val:   22,
		Left:  nil,
		Right: nil,
	}
	a2 := &TreeNode{
		Val:   2,
		Left:  a4,
		Right: nil,
	}
	a3 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: a5,
	}
	a1 := &TreeNode{
		Val:   1,
		Left:  a2,
		Right: a3,
	}

	rs := isSymmetric(a1)
	fmt.Println(rs)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	return sim(root.Left, root.Right)

}

func sim(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && sim(left.Left, right.Right) && sim(left.Right, right.Left)
}
