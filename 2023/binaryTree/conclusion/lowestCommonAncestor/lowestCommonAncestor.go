package main

import "fmt"

func main() {
	r7 := &TreeNode{
		Val: 7,
	}
	l6 := &TreeNode{
		Val: 6,
	}
	r5 := &TreeNode{
		Val: 5,
	}
	l4 := &TreeNode{
		Val: 4,
	}
	r3 := &TreeNode{
		Val:   3,
		Left:  l6,
		Right: r7,
	}
	l2 := &TreeNode{
		Val:   2,
		Left:  l4,
		Right: r5,
	}
	root1 := &TreeNode{
		Val:   1,
		Left:  l2,
		Right: r3,
	}
	r := lowestCommonAncestor(root1, r7, l6)
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}

	return right
}
